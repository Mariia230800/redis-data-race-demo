package app

import (
	"context"
	"fmt"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/config"
	"github.com/Mariia230800/redis-data-race-demo/internal/ifrastructure/redis"
	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/repository"
	cache "github.com/Mariia230800/redis-data-race-demo/internal/repository/redis"
	"github.com/Mariia230800/redis-data-race-demo/internal/service"
)

func Run(ctx context.Context) error {
	cfg := config.Load()
	log.Init(cfg.Logger.Level)
	logger := log.Get()
	logger.Infof("Logger initialized with level: %s", cfg.Logger.Level)

	redisClient, err := redis.InitRedis(cfg)
	if err != nil {
		return fmt.Errorf("redis init: %w", err)
	}

	ttl := time.Duration(cfg.Redis.CacheTTLHours) * time.Hour
	cache := cache.NewRedisCache(redisClient, ttl) // ttl задаёт время жизни кеша (24 часа)

	db := repository.NewMockRepo()

	service := service.NewService(db, cache)

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
