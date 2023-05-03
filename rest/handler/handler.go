package handler

import (
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/grpc/client"
	"github/achjailani/go-simple-grpc/pkg/logger"
)

// Handler is a struct
type Handler struct {
	client *client.GRPCClient
	repo   *service.Repositories
	logger *logger.Logger
}

// NewHandler is a function
func NewHandler(repo *service.Repositories, client *client.GRPCClient, loggr *logger.Logger) *Handler {
	return &Handler{
		repo:   repo,
		client: client,
		logger: loggr,
	}
}
