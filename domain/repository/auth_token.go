package repository

import "github/achjailani/go-simple-grpc/domain/entity"

// AuthTokenRepoInterface is an interface repo
type AuthTokenRepoInterface interface {
	CreateAuthToken(*entity.AuthToken, *entity.User) (*entity.AuthToken, error)
	GetAuthToken(int) (*entity.AuthToken, error)
	GetAuthTokenByToken(string) (*entity.AuthToken, error)
}
