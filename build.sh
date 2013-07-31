#!/bin/bash

mkdir -p bin

cd main
go build -o MrGo master.go
go build -o MrWorker worker.go
mv MrGo ../bin
mv MrWorker ../bin
cd ..
