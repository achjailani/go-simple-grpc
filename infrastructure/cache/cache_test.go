package cache_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/infrastructure/cache"
	"github/achjailani/go-simple-grpc/tests"
	"testing"
	"time"
)

func TestNew_WithRedis(t *testing.T) {
	test := tests.Init()

	ctx := test.Ctx

	redis, err := cache.NewRedisCache(test.Cfg)

	assert.NoError(t, err)
	assert.NotNil(t, redis)

	cacheInstance := cache.New(redis)

	assert.NotNil(t, cacheInstance)
	fmt.Println(test.Cfg.CacheConfig)

	key := "greeting"
	val := "Hi, there!"

	t.Run("it should valid set", func(t *testing.T) {
		errSet := cacheInstance.Set(ctx, key, val, 1*time.Second)

		assert.NoError(t, errSet)
	})

	t.Run("it should valid get", func(t *testing.T) {
		r, errGet := cacheInstance.Get(ctx, key)

		assert.NoError(t, errGet)
		assert.NotNil(t, r)
		assert.Equal(t, val, r.(string))
	})
}
