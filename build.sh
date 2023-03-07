#!/bin/sh

mkdir -p bin
cd server
go build -o ../bin/server
cd ../
echo 'Done'