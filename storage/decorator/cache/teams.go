package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/serhiihuberniuk/unstable-api/models"
)

func (f *CacheDecorator) Teams(ctx context.Context) ([]models.Team, error) {
	data, ok := f.cache.Get("/teams")
	if ok {
		teams, ok := data.([]models.Team)
		if !ok {
			return nil, errors.New("error: cannot assert data from cache to []models.Team")
		}

		return teams, nil
	}

	teams, err := f.fetcher.Teams(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching data from remote: %w", err)
	}

	f.cache.Put(ctx, "/teams", teams)

	return teams, nil
}
