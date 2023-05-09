package datalayer

import (
	"fmt"

	"github.com/luisnquin/blind-creator-rest-api-test/internal/config"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/log"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewForORM(config config.App) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		config.Database.Host(),
		config.Database.User(),
		config.Database.Password(),
		config.Database.Name(),
		config.Database.Port(),
	)

	log.Trace().Msg("connecting to database...")

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	log.Trace().Msg("successfully connected...")

	log.Trace().Msg("ensuring database tables...")

	err = db.AutoMigrate(&models.Customer{}, &models.Order{}, &models.OrderItem{}, &models.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
