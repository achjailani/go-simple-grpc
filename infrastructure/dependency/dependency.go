package dependency

import (
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/grpc/client"
	"github/achjailani/go-simple-grpc/infrastructure/cache"
	"github/achjailani/go-simple-grpc/pkg/logger"
)

// Dependency is a struct
type Dependency struct {
	Cfg        *config.Config
	Repo       *service.Repositories
	Logger     *logger.Logger
	GRPCClient *client.GRPCClient
	Cache      *cache.Cache
}

// New is a constructor
func New(opts ...Option) *Dependency {
	dep := &Dependency{}

	for _, opt := range opts {
		opt.apply(dep)
	}

	return dep
}
