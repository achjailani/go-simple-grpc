package handler

import (
	"github/achjailani/go-simple-grpc/client"
	"github/achjailani/go-simple-grpc/domain/service"
)

// Handler is a struct
type Handler struct {
	client *client.GRPCClient
	repo   *service.Repositories
}

// NewHandler is a function
func NewHandler(repo *service.Repositories, client *client.GRPCClient) *Handler {
	return &Handler{
		repo:   repo,
		client: client,
	}
}
