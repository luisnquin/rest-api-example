#!/bin/bash

if command -v pg-ping &>/dev/null; then
    echo "waiting for database mount..."
    pg-ping -U "${POSTGRES_USER}" -p "${POSTGRES_PASSWORD}" -d "${POSTGRES_DB}" -h "${POSTGRES_HOST}:${POSTGRES_PORT}" --exit-on-success
else
    echo "waiting for 5 seconds..."
    sleep 5
fi
