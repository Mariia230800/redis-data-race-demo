package app

import (
	"context"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/config"
	"github.com/Mariia230800/redis-data-race-demo/internal/log"
)

func Run(ctx context.Context) error {
	cfg := config.Load()
	log.Init(cfg.Logger.Level)
	logger := log.Get()
	logger.Infof("Logger initialized with level: %s", cfg.Logger.Level)

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		close(done)
	}()

	logger.Info("Cron-service running...")
	for {
		logger.Info("Cron heartbeat")
		time.Sleep(30 * time.Second)
	}
}
