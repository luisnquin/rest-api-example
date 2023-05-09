package config

import (
	"fmt"
	"strconv"
)

type Server struct{}

func (s Server) Port() string {
	value := mustEnv("PORT")

	if value[0] != ':' {
		value = ":" + value
	}

	if _, err := strconv.Atoi(value[1:]); err != nil {
		panic(fmt.Sprintf("invalid PORT value: %s", value))
	}

	return value
}
