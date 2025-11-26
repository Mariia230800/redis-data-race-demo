package redis

import (
	"encoding/json"
	"fmt"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/models"
)

func (c *RedisCache) Set(movie models.Movie) error {
	key := fmt.Sprintf("movie:%s", movie.ID)
	data, err := json.Marshal(movie)
	if err != nil {
		log.Errorf("Failed to marshal movie %s: %v", movie.ID, err)
		return fmt.Errorf("marshal failed for movie %s: %w", movie.ID, err)
	}
	if err := c.client.Set(c.ctx, key, data, c.ttl).Err(); err != nil {
		log.Errorf("Redis SET failed for movie %s: %v", movie.ID, err)
		return fmt.Errorf("redis set failed for movie %s: %w", movie.ID, err)
	}
	log.Infof("Redis SET success for movie %s", movie.ID)
	return nil
}
