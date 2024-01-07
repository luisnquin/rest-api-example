package database

import (
	"fmt"

	"github.com/luisnquin/server-example/internal/config"
)

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
