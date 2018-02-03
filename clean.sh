#!/bin/bash
echo "cleaning..."
docker rmi -f tester
rm -rf ./bin
