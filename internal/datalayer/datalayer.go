package datalayer

import (
	"fmt"
	"time"

	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creates a new connection (via ORM) using the provided database parameters that were specified in the configuration.
func Connect(config config.App) (*gorm.DB, error) {
	log.Trace().Msg("connecting to database...")

	dsn := generateDsnFromConfig(config)
	dialect := postgres.Open(dsn)

	retryInterval := time.Second * 2
	maxRetries := uint8(5)

	db, err := connectORM(dialect, maxRetries, 0, retryInterval)
	if err != nil {
		return nil, err
	}

	log.Trace().Msg("successfully connected...")

	return db, nil
}

func connectORM(dialect gorm.Dialector, maxRetries, attempts uint8, retryInterval time.Duration) (*gorm.DB, error) {
	db, err := gorm.Open(dialect)
	if err != nil {
		if getSQLErrorCode(err) == cannot_connect_now_code && attempts < maxRetries {
			log.Info().Msgf("database is not reachable yet, retrying in %s seconds...", retryInterval)
			time.Sleep(retryInterval)

			return connectORM(dialect, maxRetries, attempts+1, retryInterval+time.Second*5)
		}

		return nil, err
	}

	return db, nil
}

func MigrateUsingORM(db *gorm.DB) error {
	log.Trace().Msg("ensuring database tables...")

	return db.AutoMigrate(&models.Customer{}, &models.Order{}, &models.OrderItem{}, &models.Product{})
}

func DropAllUsingORM(db *gorm.DB) error {
	log.Trace().Msg("dropping all database tables...")

	return db.Migrator().DropTable(&models.Customer{}, &models.Order{}, &models.OrderItem{}, &models.Product{})
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
