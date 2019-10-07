#!/usr/bin/env bash

CONTAINER="pg-dev"

docker exec -it ${CONTAINER} psql -U postgres -c "DROP DATABASE IF EXISTS authdb"
docker exec -it ${CONTAINER} psql -U postgres -c "CREATE DATABASE authdb"

# Create tables
docker exec -it ${CONTAINER} psql -U postgres -c "\i /postgres/sql/schema.pgsql"

# Create test data
docker exec -it ${CONTAINER} psql -U postgres -c "\i /postgres/sql/populate_security_level.pgsql"
docker exec -it ${CONTAINER} psql -U postgres -c "\i /postgres/sql/populate_position.pgsql"
docker exec -it ${CONTAINER} psql -U postgres -c "\i /postgres/sql/populate_fg_user.pgsql"
docker exec -it ${CONTAINER} psql -U postgres -c "\i /postgres/sql/populate_user_position_relation.pgsql"



