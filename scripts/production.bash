#!/usr/bin/env bash
docker-compose -f docker-compose.core.yml -f docker-compose.prod.yml up -d --build
./postgres/migrate.prod.bash

