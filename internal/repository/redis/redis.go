package redis

import (
	"context"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/models"
	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Set(movie models.Movie) error
	Get(id string) (*models.Movie, error)
	GetMovies(ctx context.Context) ([]models.Movie, error)
	SetMovies(ctx context.Context, movies []models.Movie) error
	Ping() error
}

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
	ctx    context.Context
}

func NewRedisCache(client *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{
		client: client,
		ttl:    ttl,
	}
}

func (c *RedisCache) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	status := c.client.Ping(ctx)
	return status.Err()
}
