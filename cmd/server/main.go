package main

import (
	"context"

	"github.com/luisnquin/server-example/internal/business/locations"
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/database"
	"github.com/luisnquin/server-example/internal/database/sqlc"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/server"
)

func main() {
	if err := showDebugInfo(); err != nil {
		panic(err)
	}

	appConfig := config.NewApp()

	server := server.New(appConfig)

	ctx := context.Background()

	pool, err := database.NewConnectionPool(ctx, appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create connection with db")
	}

	defer pool.Close()

	querier := sqlc.New(pool)

	locations.NewManager(querier).RegisterHandlers(server)

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("while the server was running...")
	}
}
