#!/bin/sh

BASEDIR=$(pwd)
#mkdir -p bin
cd cmd/pengha-api
env GOOS=linux GOARCH=arm go build -o ../../bin/pengha-api
cp ./conf/config.yml ../../bin/

cd $BASEDIR
VERSION=$(cat version.txt)
docker build -t yonghochoi/pengha-api:$VERSION .
docker push yonghochoi/pengha-api:$VERSION
