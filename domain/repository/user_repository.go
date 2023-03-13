package repository

import (
	"context"
	"github/achjailani/go-simple-grpc/domain/entity"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, id int, user *entity.User) error
	Find(ctx context.Context, id int) (*entity.User, error)
	Get(ctx context.Context) ([]entity.User, error)
	Delete(ctx context.Context, id int) error
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
}
