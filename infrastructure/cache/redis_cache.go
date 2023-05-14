package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github/achjailani/go-simple-grpc/config"
	"time"
)

// RedisCache is a type
type RedisCache struct {
	cfg    *config.Config
	client *redis.Client
}

// Get is a method
func (r *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	return r.client.Get(ctx, key).Result()
}

// Set is a method
func (r *RedisCache) Set(ctx context.Context, key string, val interface{}, exp time.Duration) error {
	return r.client.Set(ctx, key, val, exp).Err()
}

// Delete is a method
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// NewRedisCache is a constructor
func NewRedisCache(cfg *config.Config) (*RedisCache, error) {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.CacheConfig.CacheHost, cfg.CacheConfig.CachePort),
		Password: cfg.CacheConfig.CachePassword,
		DB:       cfg.CacheConfig.CacheDB,
	}

	if cfg.TestMode {
		opt = &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.TestCacheConfig.CacheHost, cfg.TestCacheConfig.CachePort),
			Password: cfg.TestCacheConfig.CachePassword,
			DB:       cfg.TestCacheConfig.CacheDB,
		}
	}

	conn := redis.NewClient(opt)

	return &RedisCache{
		cfg:    cfg,
		client: conn,
	}, nil
}

var _ CacheStrategy = &RedisCache{}
