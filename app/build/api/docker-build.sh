#!/bin/bash -e
cd `dirname $0`/../../

docker build -t $ENV/api -f ./build/api/Dockerfile .

cd -
