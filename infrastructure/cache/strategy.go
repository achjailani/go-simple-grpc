package cache

import (
	"context"
	"time"
)

// CacheStrategy is an interface
type CacheStrategy interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, val interface{}, exp time.Duration) error
	Delete(ctx context.Context, key string) error
}
