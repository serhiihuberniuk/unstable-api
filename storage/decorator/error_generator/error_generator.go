package error_generator

import (
	"context"
	"math/rand"

	"github.com/serhiihuberniuk/unstable-api/models"
)

type decorated interface {
	Leagues(ctx context.Context) ([]models.Leagues, error)
	Teams(ctx context.Context) ([]models.Team, error)
}

type ErrorGeneratorDecorator struct {
	decorated decorated
}

func NewErrorGeneratorDecorator(d decorated) *ErrorGeneratorDecorator {
	return &ErrorGeneratorDecorator{
		decorated: d,
	}
}

func generateError() error {
	if rand.Float64() < 0.2 {
		if rand.Float64() >= 0.5 {
			return models.ErrServer
		}

		return models.ErrTimeout
	}

	return nil
}
