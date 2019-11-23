#!/bin/sh

BASEDIR=$(pwd)
cd cmd/pengha-api
env GOOS=linux GOARCH=amd64 go build -o ../../bin/pengha-api
cp ./conf/config.yml ../../bin/

cd $BASEDIR
VERSION=$(cat version.txt)
docker build -t pengha-api .
docker tag pengha-api yonghochoi/pengha-api:$VERSION
docker tag pengha-api yonghochoi/pengha-api:latest
docker push yonghochoi/pengha-api:$VERSION
docker push yonghochoi/pengha-api:latest

docker build -t pengha-nginx -f Dockerfile_nginx .
docker tag pengha-nginx yonghochoi/pengha-nginx:latest
docker push yonghochoi/pengha-nginx:latest