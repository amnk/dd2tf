#!/bin/bash

declare -r IMAGE_NAME="amnk/dd2tf"
declare -r IMAGE_TAG="latest"
declare -r DATADOG_API_KEY=xxx
declare -r DATADOG_APP_KEY=xxx

echo "Starting container for image '$IMAGE_NAME:$IMAGE_TAG'"
docker run -e DATADOG_API_KEY=$DATADOG_API_KEY -e DATADOG_APP_KEY=$DATADOG_APP_KEY -v ${PWD}/exports:/app/exports $IMAGE_NAME:$IMAGE_TAG monitors --ids xxx

