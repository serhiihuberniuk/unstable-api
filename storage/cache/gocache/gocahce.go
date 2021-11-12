package gocache

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
)

type GoCache struct {
	cache *cache.Cache
}

func New(defaultExpiration time.Duration, cleanupInterval time.Duration) *GoCache {
	c := cache.New(defaultExpiration, cleanupInterval)
	return &GoCache{
		cache: c,
	}
}

func (c *GoCache) Put(_ context.Context, key string, value interface{}) {
	c.cache.Set(key, value, cache.DefaultExpiration)
}

func (c *GoCache) Get(key string) (interface{}, bool) {
	data, ok := c.cache.Get(key)
	if ok {
		return data, true
	}

	return nil, false
}
