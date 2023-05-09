package config

import "os"

type Database struct{}

func (d Database) Port() string {
	return os.Getenv("POSTGRES_PORT")
}

func (d Database) Name() string {
	return os.Getenv("POSTGRES_DB")
}

func (d Database) Password() string {
	return os.Getenv("POSTGRES_PASSWORD")
}

func (d Database) User() string {
	return os.Getenv("POSTGRES_USER")
}
