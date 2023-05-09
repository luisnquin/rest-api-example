package main

import (
	"encoding/json"
	"fmt"

	"github.com/luisnquin/blind-creator-rest-api-test/internal/config"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/datalayer"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/log"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/repository/orders"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/server"
)

func main() {
	appConfig := config.NewApp()

	server := server.New(appConfig)

	db, err := datalayer.NewForORM(appConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to create connection with db")
	}

	orders, err := orders.NewRepository(db).PaginatedSearch(1, 100)
	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(&orders, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
