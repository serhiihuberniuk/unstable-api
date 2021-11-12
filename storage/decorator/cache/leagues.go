package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func (f *CacheDecorator) Leagues(ctx context.Context) ([]models.Leagues, error) {
	data, ok := f.cache.Get("/leagues")
	if ok {
		leagues, ok := data.([]models.Leagues)
		if !ok {
			return nil, errors.New("error: cannot assert data from cache to []models.Leagues")
		}

		return leagues, nil
	}

	leagues, err := f.fetcher.Leagues(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching data from remote: %w", err)
	}

	f.cache.Put(ctx, "/leagues", leagues)

	return leagues, nil
}
