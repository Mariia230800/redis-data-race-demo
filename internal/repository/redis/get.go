package redis

import (
	"encoding/json"
	"fmt"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/models"
	"github.com/redis/go-redis/v9"
)

func (c *RedisCache) Get(id string) (*models.Movie, error) {
	key := fmt.Sprintf("movie:%s", id)
	data, err := c.client.Get(c.ctx, key).Bytes()
	if err == redis.Nil {
		log.Infof("Redis key %s not found", key)
		return nil, nil
	}
	if err != nil {
		log.Errorf("Redis GET error for movie %s: %v", id, err)
		return nil, fmt.Errorf("redis get failed for movie %s: %w", id, err)
	}
	var movie models.Movie
	if err := json.Unmarshal(data, &movie); err != nil {
		log.Errorf("Redis GET unmarshal error for movie %s: %v", id, err)
		return nil, fmt.Errorf("redis unmarshal failed for movie %s: %w", id, err)
	}
	log.Infof("Redis GET success for movie %s", id)
	return &movie, nil
}
