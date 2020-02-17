#!/bin/bash

usage() {
	echo -e "\nUsage:\n./run_tar_exports.sh [optional filename goes here] \n" 
} 

# if -h/--help is supplied, display usage message and exit
if [[ ( $# == "--help") ||  $# == "-h" ]] 
then 
    usage
    exit 1
fi

# if ${PWD}/exports doesn't exist we can't tar it, so exit
if [ ! -d "${PWD}/exports" ]; then
    echo -e "ERROR: ./exports directory doesn't exist so we can't tar it, exiting...\n"
    exit 2
fi

# set tar default filename
TAR_FILENAME="exports.tar.gz"

# if optional filename argument sent to this script, then use it as tar filename
if [ -n "${1}" ]; then
    TAR_FILENAME="${1}"
fi

# ensure any prior created $TAR_FILENAME is removed
rm -f ${PWD}/exports/${TAR_FILENAME} ${PWD}/exports/exports.tar.gz

echo "Terraform validated. Creating tar archive of exported Datadog Terraform files..."
# tar the exports/ directory
docker run --rm -v ${PWD}/exports:/app/exports -w /app/exports debian:stable tar -C /app/exports --exclude=./.* -czvf "./${TAR_FILENAME}" . 
