package route

import (
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/grpc/client"
	"github/achjailani/go-simple-grpc/pkg/logger"
)

// Router is a struct contains dependencies needed
type Router struct {
	config *config.Config
	repo   *service.Repositories
	client *client.GRPCClient
	logger *logger.Logger
}

// RouterOption return Router with RouterOption to fill up the dependencies
type RouterOption func(*Router)

// NewRouter is a constructor will initialize Router.
func NewRouter(options ...RouterOption) *Router {
	router := &Router{}

	for _, opt := range options {
		opt(router)
	}

	return router
}
