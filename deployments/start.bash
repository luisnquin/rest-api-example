#!/bin/bash

EC2_INSTANCE_CONNECTION=$1
PEM_FILE_PATH=$2

PROJECT_DIR="$(pwd)"
BUILD_DIR="$PROJECT_DIR/build"

just build
just build-migrator

ssh -i "$PEM_FILE_PATH" -o StrictHostKeyChecking=no "$EC2_INSTANCE_CONNECTION" 'sudo systemctl stop server'

scp -i "$PEM_FILE_PATH" "$BUILD_DIR/server" "$EC2_INSTANCE_CONNECTION:/home/ec2-user/server"
scp -i "$PEM_FILE_PATH" "$BUILD_DIR/migrator" "$EC2_INSTANCE_CONNECTION:/home/ec2-user/migrator"
scp -i "$PEM_FILE_PATH" "$PROJECT_DIR/deployments/.env.production" "$EC2_INSTANCE_CONNECTION:/home/ec2-user/.env"
scp -i "$PEM_FILE_PATH" "$PROJECT_DIR/deployments/server.service" "$EC2_INSTANCE_CONNECTION:/tmp/server.service"

ssh -i "$PEM_FILE_PATH" -o StrictHostKeyChecking=no "$EC2_INSTANCE_CONNECTION" 'export $(xargs <.env) && ./migrator'
ssh -i "$PEM_FILE_PATH" -o StrictHostKeyChecking=no "$EC2_INSTANCE_CONNECTION" "sudo mv /tmp/server.service /etc/systemd/system/"

ssh -i "$PEM_FILE_PATH" -o StrictHostKeyChecking=no "$EC2_INSTANCE_CONNECTION" "sudo systemctl enable server && sudo systemctl start server"
