package service

import (
	"context"
	"fmt"

	"github.com/Mariia230800/redis-data-race-demo/internal/log"
	"github.com/Mariia230800/redis-data-race-demo/internal/models"
	"github.com/Mariia230800/redis-data-race-demo/internal/repository"
	"github.com/Mariia230800/redis-data-race-demo/internal/repository/redis"
)

type Servicer interface {
	GetMovies(ctx context.Context) ([]models.Movie, error)
}

type Service struct {
	repo  repository.MoviesRepository
	cache redis.RedisRepository
}

func NewService(repo repository.MoviesRepository, cache redis.RedisRepository) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

func (s *Service) GetMovies(ctx context.Context) ([]models.Movie, error) {

	cached, err := s.cache.GetMovies(ctx)
	if err == nil && len(cached) > 0 {
		return filterByYear(cached), nil
	}

	movies, err := s.repo.GetMovies(ctx)
	if err != nil {
		log.Errorf("error getting movies: %v", err)
		return nil, fmt.Errorf("failed to get list of movies: %w", err)

	}

	filtred := filterByYear(movies)

	_ = s.cache.SetMovies(ctx, filtred)

	return filtred, nil
}

func filterByYear(movies []models.Movie) []models.Movie {
	var res []models.Movie

	for _, movie := range movies {
		if movie.Year >= 2000 {
			res = append(res, movie)
		}
	}
	return res
}
