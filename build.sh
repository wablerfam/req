#!/bin/sh

if [ "${GITHUB_TOKEN}" = "" ]; then
    echo "Error: GITHUB_TOKEN is not set"
    exit 1
fi

if [ $# -ne 1 ]; then
    echo "Error: Invalid argument"
    exit 1
fi

goxc -pv=${1} \
    -arch="386 amd64" \
    -os="linux windows" \
    -+tasks=clean,compile,archive \
    publish-github \
        -owner=wablerfam \
        -repository=req \
        -apikey=${GITHUB_TOKEN}\
        -include="*"