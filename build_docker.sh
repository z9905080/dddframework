#!/bin/bash

eval "$(grep ^version_id= .env)"
docker build -t z9905080/ddd-framework:$version_id . --no-cache