package cache

import (
	"context"
	"time"
)

const (
	expiration = 24 * time.Hour
)

// Cache is a type
type Cache struct {
	strategy CacheStrategy
}

// New is a constructor
func New(strategy CacheStrategy) *Cache {
	cache := &Cache{
		strategy: strategy,
	}

	return cache
}

// Set is a method
func (c *Cache) Set(ctx context.Context, key string, val interface{}, duration time.Duration) error {
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
