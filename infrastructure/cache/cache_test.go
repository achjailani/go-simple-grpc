package cache_test

import (
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/infrastructure/cache"
	"testing"
)

func TestNew(t *testing.T) {
	_, _ = cache.New(config.New())
}
