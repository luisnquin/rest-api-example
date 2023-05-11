package config

import (
	"os"
	"strings"
)

type Database struct{}

func (d Database) Host() string {
	return mustEnv("POSTGRES_HOST")
}

func (d Database) Port() string {
	return strings.TrimSpace(os.Getenv("POSTGRES_PORT"))
}

func (d Database) Name() string {
	return mustEnv("POSTGRES_DB")
}

func (d Database) Password() string {
	return mustEnv("POSTGRES_PASSWORD")
}

func (d Database) User() string {
	return mustEnv("POSTGRES_USER")
}
