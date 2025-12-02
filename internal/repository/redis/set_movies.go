package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/models"
)

const lockKey = "lock:movies" //Redis ключ для распределённого лока

func (c *RedisCache) SetMovies(ctx context.Context, movies []models.Movie) error {

	ok, err := c.client.SetNX(ctx, lockKey, "1", c.ttl).Result()
	if err != nil {
		log.Errorf("Redis SETNX lock failed: %v", err)
		return fmt.Errorf("acquire redis lock: %w", err)
	}

	if !ok {
		// Лок уже захвачен другим инстансом
		log.Warnf("Redis lock for %s is already acquired, skipping write", lockKey)
		return fmt.Errorf("redis lock busy: %s", lockKey)
	}

	// Лок захвачен — имитируем долгую запись, чтобы другие инстансы столкнулись с busy
	time.Sleep(3 * time.Second)

	defer func() {
		if _, err := c.client.Del(ctx, lockKey).Result(); err != nil {
			log.Errorf("failed to release redis lock %s: %v", lockKey, err)
		}
	}()

	data, err := json.Marshal(movies)
	if err != nil {
		log.Errorf("failed to marshal movies %+v: %v", movies, err)
		return fmt.Errorf("marshal movies: %w", err)
	}

	if err := c.client.Set(ctx, "movies", data, c.ttl).Err(); err != nil {
		log.Errorf("Redis SET failed for movies %+v: %v", movies, err)
		return fmt.Errorf("redis set movies: %w", err)
	}

	return nil
}
