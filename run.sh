#!/usr/bin/bash

docker rm rinha-backend-go-db-1
docker build --tag api .
docker compose up
