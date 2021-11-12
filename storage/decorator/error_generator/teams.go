package error_generator

import (
	"context"
	"fmt"

	"github.com/serhiihuberniuk/unstable-api/models"
)

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
