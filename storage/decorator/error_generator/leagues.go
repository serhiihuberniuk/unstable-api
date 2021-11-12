package error_generator

import (
	"context"
	"fmt"

	"github.com/serhiihuberniuk/unstable-api/models"
)

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
