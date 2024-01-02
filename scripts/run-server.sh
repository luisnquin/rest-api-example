#!/bin/sh

PROJECT_NAME="server-example"

main() {
    current_project_name="$(basename "$(pwd)")"

    if [ "$current_project_name" != "$PROJECT_NAME" ]; then
        echo "folder name doesn't match project name, want '$PROJECT_NAME' but got '$current_project_name'"
        return 1
    fi

    if ! test -f ./.env; then
        echo "missing .env file!"
        return 1
    fi

    last_compose_up_ts_file_path="/tmp/$PROJECT_NAME-compose-up.txt"

    last_compose_up_dt=$(cat /tmp/$PROJECT_NAME-compose-up.txt 2>/dev/null)
    env_file_modification_dt=$(date -r .env +%s%N | cut -b1-13)

    set -e

    if [ ! "$(docker compose ls --format=json | jq ". | map(select(.Name == \"$PROJECT_NAME\")) | any")" = "true" ]; then
        echo "docker compose for current '$PROJECT_NAME' project is not running, 'just compose-up'..."
        date +%s%N | cut -b1-13 >"$last_compose_up_ts_file_path"
        just compose-up
        echo
    elif [ "$env_file_modification_dt" -gt "$last_compose_up_dt" ]; then
        echo "modified since last compose-up, 'just compose-up'..."
        just compose-up
        date +%s%N | cut -b1-13 >"$last_compose_up_ts_file_path"
        echo
    fi

    just build && just run
}

exit "$(main "$@")"
