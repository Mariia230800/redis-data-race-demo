package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/models"
	"github.com/redis/go-redis/v9"
)

func (c *RedisCache) GetMovies(ctx context.Context) ([]models.Movie, error) {
	data, err := c.client.Get(ctx, "movies").Bytes()
	if err != nil {
		if err == redis.Nil {
			log.Infof("movies cache miss")
			return nil, nil
		}
		log.Errorf("redis GET failed for movies: %v", err)
		return nil, fmt.Errorf("redis get movies: %w", err)
	}

	var movies []models.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		log.Errorf("failed to unmarshal cached movies: %v", err)
		return nil, fmt.Errorf("unmarshal movies: %w", err)
	}

	log.Infof("movies cache hit: %+v", movies)

	return movies, nil
}
