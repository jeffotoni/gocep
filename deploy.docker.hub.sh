#!/bin/bash
# Go Api server
# @jeffotoni
# 2019-06-01

echo "-------------------------------------- Clean <none> images ---------------------------------------"
docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force

echo "\033[0;33m################################## go build gocep ##################################\033[0m"
GOOS=linux go build -ldflags="-s -w" -o gocep main.go
upx gocep

echo "\033[0;33m################################## build docker gocep ##################################\033[0m"
docker build -f Dockerfile -t jeffotoni/gocep .

echo "\033[0;33m######################################### login aws and push ########################################\033[0m"
docker login
docker push jeffotoni/gocep
echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"