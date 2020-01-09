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

declare -r IMAGE_NAME="toozej/dd2tf"
declare -r IMAGE_TAG="latest"

echo "Starting container for image '$IMAGE_NAME:$IMAGE_TAG'"

# run the docker container, passing any additional arguments to ./run_dd2tf.sh as arguments to the container and thus dd2tf binary
docker run -e DATADOG_API_KEY=$DATADOG_API_KEY -e DATADOG_APP_KEY=$DATADOG_APP_KEY -v ${PWD}/exports:/app/exports $IMAGE_NAME:$IMAGE_TAG $@

