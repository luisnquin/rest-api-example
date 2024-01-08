package main

import (
	"context"
	"runtime"
	"runtime/debug"

	"github.com/luisnquin/server-example/internal/business/locations"
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/database"
	"github.com/luisnquin/server-example/internal/database/sqlc"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/server"
)

func main() {

	appConfig := config.NewApp()

	server := server.New(appConfig)
	showDebugInfo()

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

func showDebugInfo() {
	buildInfo, _ := debug.ReadBuildInfo()

	log.Trace().Str("go_version", buildInfo.GoVersion).Str("GOOS", runtime.GOOS).Str("GOARCH", runtime.GOARCH).
		Int("cpu_count", runtime.NumCPU()).Int("pid", os.Getpid()).Send()
}
