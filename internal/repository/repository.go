package repository

import (
	"context"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/models"
)

type MoviesRepository interface {
	GetMovies(ctx context.Context) ([]models.Movie, error)
}

type mockRepo struct{}

func NewMockRepo() MoviesRepository {
	return &mockRepo{}
}

func (r *mockRepo) GetMovies(ctx context.Context) ([]models.Movie, error) {
	now := time.Now().Year()

	all := []models.Movie{
		{ID: "1", Title: "The Matrix", Year: 1999},
		{ID: "2", Title: "Gladiator", Year: 2000},
		{ID: "3", Title: "Inception", Year: 2010},
		{ID: "4", Title: "The Social Network", Year: 2010},
		{ID: "5", Title: "Future Movie", Year: now},
	}

	return all, nil
}
