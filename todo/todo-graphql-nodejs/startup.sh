#!/bin/sh

if [ "$TODO_SERVICE_URL" == "" ]; then
    echo "TODO_SERVICE_URL env is required"
    exit 1
fi

node index.js
