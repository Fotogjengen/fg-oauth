#!/usr/bin/env bash

ssh -i ./id_rsa hilfling@$SERVER_IP_ADDRESS '
cd /var/www/hilfling-oauth &&
git fetch --all && 
git reset --hard origin/master &&
docker-compose -f ./docker-compose.yml -f ./docker-compose.deploy.yml up -d --build &&
echo "DONE"
'
