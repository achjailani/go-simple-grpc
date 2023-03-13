package persistence

import (
	"context"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/domain/repository"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ repository.UserRepositoryInterface = &UserRepo{}

func (u UserRepo) Create(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(&user).Error
}

func (u UserRepo) Update(ctx context.Context, id int, user *entity.User) error {
	return u.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", id).Updates(&user).Error
}

func (u UserRepo) Find(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepo) Get(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	err := u.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepo) Delete(ctx context.Context, id int) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error
}

func (u UserRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
