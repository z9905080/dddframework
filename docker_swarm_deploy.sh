#!/bin/bash
eval "$(grep ^version_id= .env)"
eval "$(grep ^ENV= .env)"
env version_id=$version_id ENV=$ENV docker stack deploy --with-registry-auth --compose-file docker-swarm-deploy.yaml ddd