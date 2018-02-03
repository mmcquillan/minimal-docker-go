#!/bin/bash
echo "building..."
mkdir -p ./bin
export CGO_ENABLED=0
go build -o "./bin/tester" -ldflags '-w -extldflags "-static"' .
docker build -t tester .
