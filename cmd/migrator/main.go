package main

import (
	_ "embed"
	"strings"

	"github.com/luisnquin/blind-creator-rest-api-test/internal/config"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/datalayer"
)

//go:embed mock.sql
var rawStatements string

func main() {
	db, err := datalayer.NewStore(config.NewApp())
	if err != nil {
		panic(err)
	}

	var statements []string

	for _, stmt := range strings.Split(rawStatements, ";") {
		if stmt != "" {
			statements = append(statements, stmt+";")
		}
	}

	for _, statement := range statements {
		if err := db.Exec(statement).Error; err != nil {
			panic(err)
		}
	}
}
