package config

import (
	"errors"

	"github.com/caarlos0/env/v8"
)

type (
	App struct {
		Server   Server
		Database Database
	}

	Server struct {
		Port string `env:"PORT"`
	}

	Database struct {
		Login    string `env:"DATABASE_LOGIN"`
		Password string `env:"DATABASE_PASSWORD"`
	}
)

func NewApp(configFilePath string) App {
	var appConfig App

	if err := env.Parse(&appConfig); err != nil {
		panic(err)
	}

	if err := appConfig.validate(); err != nil {
		panic(err)
	}

	return appConfig
}

func (a App) validate() error {
	if a.Server.Port == "" {
		return errors.New("missing 'PORT' environment variable")
	}

	return nil
}
