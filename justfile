

build:
	@go build -o ./build/server ./cmd/server/

run:
	@./build/server

start: build run
