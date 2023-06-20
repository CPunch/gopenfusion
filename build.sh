#!/bin/sh

mkdir -p bin
CGO_ENABLED=0 GOOS=linux go build -o ./bin/server
echo 'Done'