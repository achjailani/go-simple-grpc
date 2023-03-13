package handler

import (
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/proto/foo"
)

// Interface is an interface
type Interface interface {
	foo.UserServiceServer
	foo.HelloServer
	foo.AuthServer
	foo.LogServiceServer
}

// Handler is struct
type Handler struct {
	config *config.Config
	repo   *service.Repositories

	foo.UnimplementedUserServiceServer
	foo.UnimplementedHelloServer
	foo.UnimplementedAuthServer
	foo.UnimplementedLogServiceServer
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories) *Handler {
	return &Handler{
		config: conf,
		repo:   repo,
	}
}

var _ Interface = &Handler{}
