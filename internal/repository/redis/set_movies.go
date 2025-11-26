package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/models"
)

func (c *RedisCache) SetMovies(ctx context.Context, movies []models.Movie) error {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Errorf("Failed to marshal movies %+v: %v", movies, err)
		return fmt.Errorf("marshal movies: %w", err)
	}

	if err := c.client.Set(ctx, "movies", data, c.ttl).Err(); err != nil {
		log.Errorf("Redis SET failed for movies %+v: %v", movies, err)
		return fmt.Errorf("redis set movies: %w", err)
	}

	return nil
}
