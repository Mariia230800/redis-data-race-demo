package cron

import (
	"context"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/service"
)

type WriterCron struct {
	service service.Servicer
}

func NewWriterCron(service service.Servicer) *WriterCron {
	return &WriterCron{service: service}
}

func (c *WriterCron) Run(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Infof("WriterCron stopped")
			return
		case <-ticker.C:
			movies, err := c.service.GetMovies(context.Background())
			if err != nil {
				log.Errorf("WriterCron: failed to get or set movies: %v", err)
				continue
			}

			log.Infof("WriterCron: processed %d movies", len(movies))
		}
	}
}
