#!/bin/bash

declare -r IMAGE_NAME="amnk/dd2tf"
declare -r IMAGE_TAG="latest"

echo "Building image '$IMAGE_NAME:$IMAGE_TAG'..."
docker build -t $IMAGE_NAME:$IMAGE_TAG .
