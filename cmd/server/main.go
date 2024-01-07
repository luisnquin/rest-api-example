package main

import (
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/server"
)

func main() {
	appConfig := config.NewApp()

	server := server.New(appConfig)

	if err != nil {
		log.Fatal().Err(err).Msg("unable to create connection with db")
	}


	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("while the server was running...")
	}
}
