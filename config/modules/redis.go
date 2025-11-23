package modules

type Redis struct {
	RedisHost string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort string `env:"REDIS_PORT" envDefault:"6379"`
}
