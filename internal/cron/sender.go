package cron

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/config"
	"github.com/Mariia230800/redis-data-race-demo/internal/ifrastructure/kafka"
	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/service"
)

type SenderCron struct {
	service     service.Servicer
	kafkaClient kafka.Producer
	interval    time.Duration
}

func NewSenderCron(service service.Servicer, kafkaClient kafka.Producer, interval time.Duration) *SenderCron {
	return &SenderCron{
		service:     service,
		kafkaClient: kafkaClient,
		interval:    interval,
	}
}

func (c *SenderCron) Run(ctx context.Context) {
	logger := log.Get()
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Infof("SenderCron stopped")
			return
		case <-ticker.C:

			movies, err := c.service.GetMovies(ctx)
			if err != nil {
				logger.Errorf("Failed to get movies: %v", err)
				continue
			}

			if len(movies) == 0 {
				logger.Infof("No movies to send to Kafka")
				continue
			}

			data, err := json.Marshal(movies)
			if err != nil {
				logger.Errorf("Failed to marshal movies: %v", err)
				continue
			}
			cfg := config.Load()
			topic := cfg.Kafka.Topic
			key := "movies-key"
			if err := c.kafkaClient.SendMessage(topic, key, data); err != nil {
				logger.Errorf("Failed to send movies to Kafka: %v", err)
				continue
			}

			logger.Infof("Successfully sent %d movies to Kafka", len(movies))
		}
	}
}
