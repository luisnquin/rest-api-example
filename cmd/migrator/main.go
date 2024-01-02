package main

import (
	_ "embed"
	"strings"

	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/datalayer"
	"gorm.io/gorm"
)

//go:embed mock.sql
var mockStatements string

func main() {
	db, err := datalayer.NewForORM(config.NewApp())
	if err != nil {
		panic(err)
	}

	// I don't think that is necessary to handle
	// returned errors here
	datalayer.DropAllUsingORM(db)
	datalayer.MigrateUsingORM(db)

	if err := createMockData(db); err != nil {
		panic(err)
	}
}

func createMockData(db *gorm.DB) error {
	var statements []string

	for _, stmt := range strings.Split(mockStatements, ";") {
		if stmt != "" {
			statements = append(statements, stmt+";")
		}
	}

	for _, statement := range statements {
		if err := db.Exec(statement).Error; err != nil {
			return err
		}
	}

	return nil
}
