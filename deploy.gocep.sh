#!/bin/bash

DATA_ISO=$(date +%Y-%m-%d-%H-%M-%S)
echo -e "-------------------------------------- Clean <none> images ---------------------------------------"
docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force
echo -e "\033[0;33m######################################### pull ########################################\033[0m"
docker pull jeffotoni/gocep

docker-compose stop gocep
docker-compose rm --force gocep
docker-compose up -d gocep
docker-compose ps
echo -e "\033[0;32mGenerated Run docker-compose\033[0m \033[0;33m[ok]\033[0m \n"