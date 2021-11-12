//go:generate mockgen -destination=cache_mock_test.go -package=cache_test -source=cache.go
package cache

import (
	"context"

	"github.com/serhiihuberniuk/unstable-api/models"
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
	fetcher decorated
	cache   cache
}

func NewCacheDecorator(d decorated, c cache) *CacheDecorator {
	return &CacheDecorator{
		fetcher: d,
		cache:   c,
	}
}
