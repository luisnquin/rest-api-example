set dotenv-load
set export

dev:
	air

build:
	@go build -o ./build/server ./cmd/server/

run:
	@if command -v pp &> /dev/null; then ./build/server 2>&1 | pp; else ./build/server; fi

start: build run

compose-up:
	docker-compose up -d

migrate: erase-db-data compose-up
	sleep 5
	go run ./cmd/migrator/

erase-db-data:
	docker kill $(docker ps -qa) 2> /dev/null || true
	docker rm -f $(docker ps -qa) 2> /dev/null || true
	if test -d "./_data"; then sudo rm -rf ./_data/; fi

build-image:
	docker build . -t blind-creator-rest-api-test:latest
