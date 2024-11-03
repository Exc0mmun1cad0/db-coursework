#!/bin/bash

source ./cmd/env.sh

docker compose up -d
until docker exec $CONTAINER_NAME pg_isready -U $USER; do
    sleep 2
done

go run cmd/main.go

go_program_pid=$!
wait $go_program_pid

docker exec -t $CONTAINER_NAME pg_dump -U $USER $DB_NAME > dump.sql
docker compose down
