#!/bin/bash

# General variables
SCRIPTDIR=$(dirname $(realpath $0))
BASEDIR=$(dirname $SCRIPTDIR)

# Docker variables
CONTAINER_NAME=singullar_db
VOLUME_NAME=singullar_pg_volume
# Environment variables speficic for this operation : Should only load if this command is called 
ENVPATH="$BASEDIR/pkg/pmi/postgres.env"
source $ENVPATH

echo Executing: Running postgres container
echo User: $POSTGRES_USER
echo Password: $POSTGRES_PASSWORD
echo Container Name: singullar_db
echo Image: Postgres:17.5
echo With volume: $VOLUME_NAME
docker volume create "$VOLUME_NAME"
docker run \
--name "$CONTAINER_NAME" \
-v $VOLUME_NAME:/var/lib/postgresql/data \
-p 5432:5432 \
-e POSTGRES_USER="$POSTGRES_USER" \
-e POSTGRES_PASSWORD="$POSTGRES_PASSWORD" \
-d postgres:17.5