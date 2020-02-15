#!/bin/bash

declare -r DD2TF_IMAGE_NAME="toozej/dd2tf"
declare -r DD2TF_IMAGE_TAG="latest"

echo "Building image '$DD2TF_IMAGE_NAME:$DD2TF_IMAGE_TAG'..."
docker build -f Dockerfile_dd2tf -t $DD2TF_IMAGE_NAME:$DD2TF_IMAGE_TAG .
