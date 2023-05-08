package main

import (
	"github.com/luisnquin/blind-creator-rest-api-test/internal/config"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/server"
)

func main() {
	appConfig := config.NewApp("")

	server := server.New(appConfig)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
