#!/bin/sh

GOOS=linux
GOARCH=amd64

BASEDIR=$(pwd)
#mkdir -p bin
cd cmd/peng-api
go build -o ../../bin/pengha-api

cd $BASEDIR
