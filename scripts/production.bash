#!/usr/bin/env bash
PG_CONTAINER="pg-prod"

if docker-compose up -d --build ${PG_CONTAINER}; then
  ./postgres/migrate.prod.bash
  docker-compose up -d --build prod
fi

