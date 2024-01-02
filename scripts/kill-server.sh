#!/bin/sh

if [ -z "$SERVER_NAME" ]; then
    echo "Please specify the 'SERVER_NAME' environment variable."
    exit 1
fi

child_name="$SERVER_NAME"
possible_parent_name="air"

child_pid=$(pgrep -o "$child_name")
set -e

if [ -n "$child_pid" ]; then
    parent_pid=$(ps -o ppid= -p "$child_pid" | awk '{$1=$1;print}')

    kill "$child_pid"

    if [ -n "$parent_pid" ]; then
        parent_name=$(ps -o comm= -p "$parent_pid" | awk '{$1=$1;print}')

        if [ "$parent_name" = "$possible_parent_name" ]; then
            kill "$parent_pid"
        fi
    fi

else
    echo "Child process with name '$child_name' not found."
fi
