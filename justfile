set dotenv-load
set export

build:
	@go build -o ./build/server ./cmd/server/

run:
	@./build/server

start: build run

erase-db-data:
	docker kill $(docker ps -qa) 2> /dev/null || true
	docker rm -f $(docker ps -qa) 2> /dev/null || true
	sudo rm -rf ./_data/
