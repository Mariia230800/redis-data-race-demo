package modules

import "time"

type Crons struct {
	CronInterval time.Duration `env:"CRON_INTERVAL" envDefault:"30s"`
}
