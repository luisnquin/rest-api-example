package main

import (
	"context"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/gookit/color"
	"github.com/luisnquin/server-example/internal/business/locations"
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/database"
	"github.com/luisnquin/server-example/internal/database/sqlc"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/server"
)

func main() {
	showDebugInfo()

	ctx := context.Background()
	appConfig := config.NewApp()

	pool, err := database.NewConnectionPool(ctx, appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create connection with db")
	}

	defer pool.Close()

	querier := sqlc.New(pool)

	server := server.New(appConfig)
	server.RegisterBatch(
		locations.NewManager(querier),
	)

	server.OnBeforeStart = func() {
		port := appConfig.Server.Port()
		log.Trace().Msgf("listening on port %s", port)
		log.Trace().Msgf("%s curl -sH GET http://localhost%s", color.HEX("#46e0de").Sprint("try:"), port)
	}

	if err = server.Start(); err != nil {
		log.Fatal().Err(err).Msg("while the server was running...")
	}
}

func showDebugInfo() {
	buildInfo, _ := debug.ReadBuildInfo()

	log.Trace().Str("go_version", buildInfo.GoVersion).Str("GOOS", runtime.GOOS).Str("GOARCH", runtime.GOARCH).
		Int("cpu_count", runtime.NumCPU()).Int("pid", os.Getpid()).Send()
}
