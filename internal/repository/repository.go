package repository

import (
	"context"
	"time"

	"github.com/Mariia230800/redis-data-race-demo/internal/models"
)

// MoviesRepository interface
type MoviesRepository interface {
	GetMoviesAfterYear(ctx context.Context, year int) ([]models.Movie, error)
}

// Mock implementation (for demo)
type mockRepo struct{}

func NewMockRepo() MoviesRepository {
	return &mockRepo{}
}

func (r *mockRepo) GetMoviesAfterYear(ctx context.Context, year int) ([]models.Movie, error) {
	now := time.Now().Year()
	// sample data
	all := []models.Movie{
		{ID: "1", Title: "The Matrix", Year: 1999},
		{ID: "2", Title: "Gladiator", Year: 2000},
		{ID: "3", Title: "Inception", Year: 2010},
		{ID: "4", Title: "The Social Network", Year: 2010},
		{ID: "5", Title: "Future Movie", Year: now},
	}
	out := make([]models.Movie, 0, len(all))
	for _, m := range all {
		if m.Year >= year {
			out = append(out, m)
		}
	}
	return out, nil
}
