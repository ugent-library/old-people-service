#!/bin/sh

if [ -z "$PGDSN" ];then
    echo "environment variable PGDSN is empty. Should contain a connection string to postgres" >&2
    exit 1
fi

tern status --conn-string "$PGDSN" || exit 1
tern migrate --conn-string "$PGDSN" || exit 1
