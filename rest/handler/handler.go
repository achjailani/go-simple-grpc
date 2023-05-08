package handler

import (
	"github/achjailani/go-simple-grpc/infrastructure/dependency"
)

// Handler is a struct
type Handler struct {
	*dependency.Dependency
}

// NewHandler is a function
func NewHandler(dep *dependency.Dependency) *Handler {
	return &Handler{
		Dependency: dep,
	}
}
