package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/log"
)

// Creates a new connection using the provided database parameters that were specified in the configuration.
func NewConnectionPool(ctx context.Context, config config.App) (*pgxpool.Pool, error) {
	log.Trace().Msg("connecting to database...")

	// retryInterval := time.Second * 2
	// maxRetries := uint8(5)

	// db, err := connectORM(dialect, maxRetries, 0, retryInterval)
	// if err != nil {
	// 	return nil, err
	// }

	dsn := generateDsnFromConfig(config)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping db connection: %w", err)
	}

	log.Trace().Msg("successfully connected...")

	return pool, nil
}

func generateDsnFromConfig(config config.App) string {
	var sslMode string

	if config.IsProduction() {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host(),
		config.Database.User(),
		config.Database.Password(),
		config.Database.Name(),
		sslMode,
	)

	if port := config.Database.Port(); port != "" {
		dsn += fmt.Sprintf(" port=%s", port)
	}

	return dsn
}
