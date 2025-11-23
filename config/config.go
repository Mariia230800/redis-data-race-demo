package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"

	"github.com/Mariia230800/redis-data-race-demo/config/modules"
	"github.com/Mariia230800/redis-data-race-demo/internal/log"
)

type Config struct {
	Redis  modules.Redis
	Cron   modules.Crons
	Kafka  modules.Kafka
	Logger modules.Logger

	UseMocks bool `env:"USE_MOCKS" envDefault:"false"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables and defaults")
	}

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Errorf("Failed to parse env variables: %v", err)
	}

	return &cfg
}
