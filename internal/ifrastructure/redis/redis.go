package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/Mariia230800/redis-data-race-demo/internal/config"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
)

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}
	log.Infof("Redis client initialized successfully")
	return client, nil
}
