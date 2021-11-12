package custom

import (
	"context"
	"sync"
	"time"
)

type cacheItem struct {
	v     interface{}
	setAt time.Time
}

type cacheStorage struct {
	mx             *sync.RWMutex
	cache          map[string]cacheItem
	expirationTime time.Duration
}

func NewCustomStorage(ctx context.Context, expirationTime, checkExpiredInterval time.Duration) *cacheStorage {
	s := &cacheStorage{
		cache:          make(map[string]cacheItem),
		expirationTime: expirationTime,
		mx:             &sync.RWMutex{},
	}

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-time.After(checkExpiredInterval):
			s.mx.Lock()
			for k, v := range s.cache {
				if s.isExpired(v) {
					delete(s.cache, k)
				}
			}
			s.mx.Unlock()
		}
	}()

	return s
}

func (cs *cacheStorage) Put(_ context.Context, key string, value interface{}) {
	cs.mx.Lock()
	defer cs.mx.Unlock()
	cs.cache[key] = cacheItem{
		v:     value,
		setAt: time.Now(),
	}
}

func (cs *cacheStorage) Get(key string) (interface{}, bool) {
	cs.mx.RLock()
	defer cs.mx.RUnlock()

	value, ok := cs.cache[key]
	if !ok {
		return nil, false
	}

	if cs.isExpired(value) {
		return nil, false
	}

	return value.v, true
}

func (cs *cacheStorage) isExpired(v cacheItem) bool {
	return time.Now().After(v.setAt.Add(cs.expirationTime))
}
