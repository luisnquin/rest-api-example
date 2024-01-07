set dotenv-load
set export

dev: kill sqlc
	@HOT_RELOAD=true ./scripts/run-server.sh

build:
	@go build -o ./build/server-example ./cmd/server/

build-migrator:
	@go build -o ./build/migrator ./cmd/migrator/

run:
	@if command -v pp &> /dev/null; then ./build/server-example 2>&1 | pp; else ./build/server-example; fi

start: sqlc
	@./scripts/run-server.sh

kill:
	@SERVER_NAME="server-example" ./scripts/kill-server.sh

compose-up:
	@docker compose -p server-example -f ./docker/docker-compose.yml up -d

sqlc:
	@sqlc generate

compose-down:
	@docker compose -p server-example -f ./docker/docker-compose.yml down

migrate: erase-db-data compose-up build-migrator
	bash ./scripts/database-wait.bash
	./build/migrator/

deploy ec2-connection pem-file_path: 
	bash ./deployments/start.bash {{ec2-connection}} {{pem-file_path}}

erase-db-data:
	docker kill $(docker ps -qa) 2> /dev/null || true
	docker rm -f $(docker ps -qa) 2> /dev/null || true
	if test -d "./_data"; then sudo rm -rf ./_data/; fi

build-image:
	docker build . -t luisnquin-server-example:latest
