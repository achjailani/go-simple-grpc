package cache

import (
	"context"
	"sync"
	"time"
)

const (
	// DefaultExpiration is a constant
	DefaultExpiration = 24 * time.Hour
)

var (
	// drivers is a collection of supported drivers
	drivers = []string{"redis", "memcached"}
)

// Cache is a type
type Cache struct {
	mtx      *sync.Mutex
	strategy CacheStrategy
}

// New is a constructor
func New(strategy CacheStrategy) *Cache {
	cache := &Cache{
		strategy: strategy,
		mtx:      &sync.Mutex{},
	}

	return cache
}

// Set is a method
func (c *Cache) Set(ctx context.Context, key string, val interface{}, duration time.Duration) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.strategy.Set(ctx, key, val, duration)
}

// Get is a method
func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	return c.strategy.Get(ctx, key)
}

// Del is a method
func (c *Cache) Del(ctx context.Context, key string) error {
	return c.strategy.Delete(ctx, key)
}
