//go:generate mockgen -destination=cache_mock_test.go -package=cache_test -source=cache.go
package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/serhiihuberniuk/unstable-api/models"
)

const (
	teamsKey   = "/teams"
	leaguesKey = "/leagues"
)

type cache interface {
	Put(ctx context.Context, key string, value interface{})
	Get(key string) (interface{}, bool)
}

type decorated interface {
	Leagues(ctx context.Context) ([]models.Leagues, error)
	Teams(ctx context.Context) ([]models.Team, error)
}

type CacheDecorator struct {
	decorated decorated
	cache     cache
}

func NewCacheDecorator(d decorated, c cache) *CacheDecorator {
	return &CacheDecorator{
		decorated: d,
		cache:     c,
	}
}

func (f *CacheDecorator) Leagues(ctx context.Context) ([]models.Leagues, error) {
	data, ok := f.cache.Get(leaguesKey)
	if ok {
		leagues, ok := data.([]models.Leagues)
		if !ok {
			return nil, errors.New("error: cannot assert type to []models.Leagues when got from cache")
		}

		return leagues, nil
	}

	leagues, err := f.decorated.Leagues(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching data from remote: %w", err)
	}

	f.cache.Put(ctx, leaguesKey, leagues)

	return leagues, nil
}

func (f *CacheDecorator) Teams(ctx context.Context) ([]models.Team, error) {
	data, ok := f.cache.Get(teamsKey)
	if ok {
		teams, ok := data.([]models.Team)
		if !ok {
			return nil, errors.New("error: cannot assert type to []models.Team when got from cache")
		}

		return teams, nil
	}

	teams, err := f.decorated.Teams(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while fetching data from remote: %w", err)
	}

	f.cache.Put(ctx, teamsKey, teams)

	return teams, nil
}
