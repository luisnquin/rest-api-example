package main

import (
	"github.com/luisnquin/blind-creator-rest-api-test/internal/config"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/datalayer"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/log"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/server"
)

func main() {
	appConfig := config.NewApp()

	server := server.New(appConfig)

	db, err := datalayer.NewForORM(appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create connection with db")
	}

	_ = db

	if err := server.Start(); err != nil {
		panic(err)
	}
}
