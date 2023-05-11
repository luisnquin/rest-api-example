package config

import (
	"os"
	"strconv"
)

type (
	App struct {
		Server   Server
		Database Database
	}
)

func NewApp() App {
	return App{}
}

func (a App) IsProduction() bool {
	production, _ := strconv.ParseBool(os.Getenv("PRODUCTION"))

	return production
}
