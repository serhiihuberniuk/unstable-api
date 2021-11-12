//go:generate mockgen -destination=handler_mock_test.go -package=handler_test -source=handler.go
package handler

import (
	"context"

	"github.com/serhiihuberniuk/unstable-api/models"
)

type Handler struct {
	fetcher fetcher
}

func New(f fetcher) *Handler {
	return &Handler{
		fetcher: f,
	}
}

type fetcher interface {
	Leagues(ctx context.Context) ([]models.Leagues, error)
	Teams(ctx context.Context) ([]models.Team, error)
}
