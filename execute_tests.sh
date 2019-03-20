#!/usr/bin/env bash
cd src
./swag i

docker-compose down
docker-compose build
docker-compose run goapp_with_test
