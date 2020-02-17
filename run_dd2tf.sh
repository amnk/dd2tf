#!/bin/bash

usage() {
	echo -e "\nUsage:\n./run_dd2tf.sh [dd2tf_arguments] \n" 
} 

# if less than two arguments supplied or -h/--help is supplied, display usage message and exit
if [ $# -le 1 ] || [[ ( $# == "--help") ||  $# == "-h" ]] 
then 
    usage
    exit 1
fi

# if the required DATADOG app and api keys aren't exported as Bash variables, display usage message and exit
if [[ -z "${DATADOG_APP_KEY}" ]] || [[ -z "${DATADOG_API_KEY}" ]]
then
    echo -e "You must export DATADOG_API_KEY and DATADOG_APP_KEY environment variables to use this image\n"
    usage
    exit 2
fi

# create exports directory for use with dd2tf --files argument
if [ ! -d "${PWD}/exports" ]; then
    mkdir ${PWD}/exports
fi

echo "Starting export of Datadog files to Terraform configs..."

# run the dd2tf docker container, passing any additional arguments to ./run_dd2tf.sh as arguments to the container and thus dd2tf binary
docker run --rm -e DATADOG_API_KEY=$DATADOG_API_KEY -e DATADOG_APP_KEY=$DATADOG_APP_KEY -v ${PWD}/exports:/app/exports toozej/dd2tf:latest $@

# if exports is empty, exit 3
if [ ! "$(ls ${PWD}/exports/*.tf)" ]; then
    echo -e "${PWD}/exports/ directory doesn't contain any .tf files. This means the dd2tf export failed. Check log messages above.\n"
    exit 3
fi


echo "Datadog files exported. Initializing Terraform..."
# initialize Terraform in the exports/ directory
docker run --rm -v ${PWD}/exports:/app/exports -w /app/exports -e DATADOG_API_KEY=$DATADOG_API_KEY -e DATADOG_APP_KEY=$DATADOG_APP_KEY hashicorp/terraform:light init

echo "Terraform initialized. Validating exported Datadog files are valid Terraform configs..."
# validate Terraform files in the exports/ directory
docker run --rm -v ${PWD}/exports:/app/exports -w /app/exports -e DATADOG_API_KEY=$DATADOG_API_KEY -e DATADOG_APP_KEY=$DATADOG_APP_KEY hashicorp/terraform:light validate
