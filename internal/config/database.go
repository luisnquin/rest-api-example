package config

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/luisnquin/server-example/internal/log"
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
	secretName, region := os.Getenv("AWS_SECRETS_POSTGRES"), os.Getenv("AWS_REGION")

	if secretName != "" && region != "" {
		config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			log.Fatal().Err(err).Msg("unable to load default configuration for region (while trying to get 'database secret')")
		}

		svc := secretsmanager.NewFromConfig(config)

		input := &secretsmanager.GetSecretValueInput{
			SecretId:     aws.String(secretName),
			VersionStage: aws.String("AWSCURRENT"),
		}

		result, err := svc.GetSecretValue(context.TODO(), input)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to get 'database secret'")
		}

		var secrets map[string]string

		if err := json.Unmarshal([]byte(*result.SecretString), &secrets); err != nil {
			log.Fatal().Err(err).Msg("cannot unmarshal secret using a map[string]string")
		}

		return secrets["POSTGRES_PASSWORD"]
	}

	return mustEnv("POSTGRES_PASSWORD")
}

func (d Database) User() string {
	return mustEnv("POSTGRES_USER")
}
