package error_generator

import (
	"context"
	"fmt"
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

func (f *ErrorGeneratorDecorator) Leagues(ctx context.Context) ([]models.Leagues, error) {
	if err := generateError(); err != nil {
		return nil, err
	}

	leagues, err := f.decorated.Leagues(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching data from decorated: %w", err)
	}

	return leagues, nil
}

func (f *ErrorGeneratorDecorator) Teams(ctx context.Context) ([]models.Team, error) {
	if err := generateError(); err != nil {
		return nil, err
	}

	teams, err := f.decorated.Teams(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching data from decorated: %w", err)
	}

	return teams, nil
}
