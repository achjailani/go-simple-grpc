package repository

import "github/achjailani/go-simple-grpc/domain/entity"

type AuthTokenRepository interface {
	CreateAuthToken(*entity.AuthToken, *entity.User) (*entity.AuthToken, error)
	GetAuthToken(int) (*entity.AuthToken, error)
	GetAuthTokenByToken(string) (*entity.AuthToken, error)
}
