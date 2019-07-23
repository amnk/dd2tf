#!/bin/bash

usage() { 
	echo -e "\nUsage:\n./run_dd2tf.sh [dd2tf_arguments] \n" 
} 

# if less than two arguments supplied or -h/--help supplied, display usage 
if [ $# -le 1 ] || [[ ( $# == "--help") ||  $# == "-h" ]] 
then 
    usage
    exit 1
fi 

declare -r IMAGE_NAME="amnk/dd2tf"
declare -r IMAGE_TAG="latest"
declare -r DATADOG_API_KEY=xxx
declare -r DATADOG_APP_KEY=xxx

echo "Starting container for image '$IMAGE_NAME:$IMAGE_TAG'"
docker run -e DATADOG_API_KEY=$DATADOG_API_KEY -e DATADOG_APP_KEY=$DATADOG_APP_KEY -v ${PWD}/exports:/app $IMAGE_NAME:$IMAGE_TAG $@

