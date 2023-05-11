set dotenv-load
set export

dev:
	air

build:
	@go build -o ./build/server ./cmd/server/

build-migrator:
	@go build -o ./build/migrator ./cmd/migrator/

run:
	@if command -v pp &> /dev/null; then ./build/server 2>&1 | pp; else ./build/server; fi

start: build run

compose-up:
	docker-compose up -d

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
	docker build . -t blind-creator-rest-api-test:latest
