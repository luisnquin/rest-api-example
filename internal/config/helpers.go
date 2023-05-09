package config

import (
	"fmt"
	"os"
	"strings"
)

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("missing environment variable %q", key))
	}

	return strings.TrimSpace(value)
}
