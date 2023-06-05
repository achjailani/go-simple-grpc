package repository

import (
	"context"
	"github/achjailani/go-simple-grpc/domain/entity"
)

type PersonalInfoRepositoryInterface interface {
	Create(ctx context.Context, personalInfo *entity.PersonalInfo) error
	Update(ctx context.Context, id int, personalInfo *entity.PersonalInfo) error
	Find(ctx context.Context, id int) (*entity.PersonalInfo, error)
	Get(ctx context.Context) ([]entity.PersonalInfo, error)
	Delete(ctx context.Context, id int) error
	FindByEmail(ctx context.Context, blindIndex string) (*entity.PersonalInfo, error)
	FindByPhone(ctx context.Context, blindIndex string) (*entity.PersonalInfo, error)
}
