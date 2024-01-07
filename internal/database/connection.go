package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/log"
)

type retryParams struct {
	Retries    int8
	MaxRetries int8
	Interval   time.Duration
}

func defaultRetryParams() retryParams {
	return retryParams{
		Interval:   time.Second * 2,
		Retries:    0,
		MaxRetries: 5,
	}
}

func (c retryParams) next() retryParams {
	c.Retries++
	c.Interval += time.Second * 5

	return c
}

// Creates a new pool connection using the provided database parameters that were specified in the configuration.
func NewConnectionPool(ctx context.Context, config config.App) (*pgxpool.Pool, error) {
	log.Trace().Msg("connecting to database...")

	dsn := generateDsnFromConfig(config)

	pool, err := newConnectionPool(ctx, dsn, defaultRetryParams())
	if err != nil {
		return nil, err
	}

	log.Trace().Msg("successfully connected...")

	return pool, nil
}

func newConnectionPool(ctx context.Context, dsn string, rp retryParams) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		if getSQLErrorCode(err) == cannot_connect_now_code && rp.Retries < rp.MaxRetries {
			log.Warn().Msgf("we tried to connect to database too early, retrying in %s...", rp.Interval)
			time.Sleep(rp.Interval)

			return newConnectionPool(ctx, dsn, rp.next())
		}

		return nil, fmt.Errorf("unable to ping db connection: %w", err)
	}

	return pool, nil
}
