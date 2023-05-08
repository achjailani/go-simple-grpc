package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github/achjailani/go-simple-grpc/config"
	"time"
)

// Cache is a type
type Cache struct {
	*redis.Client
}

// New is a constructor
func New(cfg *config.Config) (*Cache, error) {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisTestConfig.RedisHost, cfg.RedisTestConfig.RedisPort),
		Password: cfg.RedisTestConfig.RedisPassword,
		DB:       cfg.RedisTestConfig.RedisDB,
	}

	if cfg.TestMode {
		opt = &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.RedisTestConfig.RedisHost, cfg.RedisTestConfig.RedisPort),
			Password: cfg.RedisTestConfig.RedisPassword,
			DB:       cfg.RedisTestConfig.RedisDB,
		}
	}

	conn := redis.NewClient(opt)

	return &Cache{
		Client: conn,
	}, nil
}

// Set is a method
func (c *Cache) Set(key string, val string) error {
	//TODO implement me
	panic("implement me")
}

// SetX is a method
func (c *Cache) SetX(key string, val string, ttl time.Duration) error {
	//TODO implement me
	panic("implement me")
}

// Get is a method
func (c *Cache) Get(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

// Del is a method
func (c *Cache) Del() error {
	//TODO implement me
	panic("implement me")
}
