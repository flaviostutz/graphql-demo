#!/bin/sh

if [ "$FEDERATED_GRAPHQL_URLS" == "" ]; then
    echo "FEDERATED_GRAPHQL_URLS env is required"
    exit 1
fi

node index.js
