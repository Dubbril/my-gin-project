package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var (
	once      sync.Once
	envConfig *EnvConfig
)

type EnvConfig struct {
	PostgresConnection string        `env:"DB_CONNECT,notEmpty"`
	ConnectionTimeout  time.Duration `env:"CONNECTION_TIMEOUT,notEmpty"`
	ExternalUrl        string        `env:"EXTERNAL_URL,notEmpty"`
}

func GetConfig() *EnvConfig {
	once.Do(func() {
		// Mock env for develop
		err := godotenv.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Cannot load env")
		}
		envConfig = loadConfig()
	})

	return envConfig
}

func loadConfig() *EnvConfig {
	var dbConfig EnvConfig
	if err := env.Parse(&dbConfig); err != nil {
		log.Fatal().Err(err).Msg("Cannot load envConfig")
	}
	return &dbConfig
}
