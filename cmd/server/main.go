package main

import (
	"github.com/luisnquin/server-example/internal/business"
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/datalayer"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/server"
)

func main() {
	appConfig := config.NewApp()

	server := server.New(appConfig)

	db, err := datalayer.Connect(appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create connection with db")
	}

	businessMgr := business.NewManager(db)
	businessMgr.RegisterHandlers(server)

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("while the server was running...")
	}
}
