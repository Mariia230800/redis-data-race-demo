package modules

type Redis struct {
	Host          string `env:"REDIS_HOST" envDefault:"localhost"`
	Port          string `env:"REDIS_PORT" envDefault:"6379"`
	Password      string `env:"REDIS_PASSWORD" envDefault:""`
	DB            int    `env:"REDIS_DB" envDefault:"0"`
	CacheTTLHours int    `env:"CACHE_TTL_HOURS" envDefault:"24"`
}
