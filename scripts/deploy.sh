#!/usr/bin/env bash

ssh -i ./deploy_key hilfling@$SERVER_IP_ADDRESS '
cd /var/www/hilfling-oauth &&
git pull &&
docker-compose -f ../docker-compose.yml -f ../docker-compose.deploy.yml up -d --build
'
