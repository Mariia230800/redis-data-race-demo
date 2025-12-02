package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/config"
	"github.com/Mariia230800/redis-data-race-demo/internal/cron"
	"github.com/Mariia230800/redis-data-race-demo/internal/ifrastructure/kafka"
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

	brokers := strings.Split(cfg.Kafka.KafkaBroker, ",")
	kafkaClient, err := kafka.NewKafkaProducer(brokers)
	if err != nil {
		return fmt.Errorf("kafka init: %w", err)
	}
	defer kafkaClient.Close()

	writerCron := cron.NewWriterCron(service)
	logger.Infof("Starting WriterCron...")
	go func() {
		writerCron.Run(ctx)
		logger.Infof("WriterCron stopped")
	}()
	senderCron := cron.NewSenderCron(service, kafkaClient, 10*time.Minute)
	logger.Infof("Starting SenderCron...")
	go func() {
		senderCron.Run(ctx)
		logger.Infof("SenderCron stopped")
	}()

	<-ctx.Done()
	logger.Info("Cron-service stopped gracefully")
	return nil
}
