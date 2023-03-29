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

// Dependency collects dependencies needed by handler
type Dependency struct {
	Config *config.Config
	Repo   *service.Repositories
}

// CoreGRPCService collects grpc service server
type CoreGRPCService struct {
	foo.UnimplementedUserServiceServer
	foo.UnimplementedHelloServer
	foo.UnimplementedAuthServer
	foo.UnimplementedLogServiceServer
}

// Handler is struct
type Handler struct {
	Dependency *Dependency

	CoreGRPCService
}

// NewHandler is a constructor
func NewHandler(conf *config.Config, repo *service.Repositories) *Handler {
	return &Handler{
		Dependency: &Dependency{
			Config: conf,
			Repo:   repo,
		},
	}
}

var _ Interface = &Handler{}
