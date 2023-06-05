package persistence

import (
	"context"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/domain/repository"
	"github/achjailani/go-simple-grpc/pkg/cryptox"
	"gorm.io/gorm"
)

// PersonalInfoRepo is a construct
type PersonalInfoRepo struct {
	db *gorm.DB
}

// NewPersonalInfoRepository is a constructor
func NewPersonalInfoRepository(db *gorm.DB) *PersonalInfoRepo {
	return &PersonalInfoRepo{db: db}
}

var _ repository.PersonalInfoRepositoryInterface = &PersonalInfoRepo{}

// Create is a method
func (u *PersonalInfoRepo) Create(ctx context.Context, personalInfo *entity.PersonalInfo) error {
	personalInfo.NameBlindIndex, _ = cryptox.MakeBlindIndex(personalInfo.Name)
	personalInfo.EmailBlindIndex, _ = cryptox.MakeBlindIndex(personalInfo.Email)
	personalInfo.PhoneBlindIndex, _ = cryptox.MakeBlindIndex(personalInfo.Phone)

	return u.db.WithContext(ctx).Create(&personalInfo).Error
}

// Update is a method
func (u *PersonalInfoRepo) Update(ctx context.Context, id int, personalInfo *entity.PersonalInfo) error {
	personalInfo.NameBlindIndex, _ = cryptox.MakeBlindIndex(personalInfo.Name)
	personalInfo.EmailBlindIndex, _ = cryptox.MakeBlindIndex(personalInfo.Email)
	personalInfo.PhoneBlindIndex, _ = cryptox.MakeBlindIndex(personalInfo.Phone)

	return u.db.WithContext(ctx).Model(&entity.PersonalInfo{}).Where("id = ?", id).Updates(personalInfo).Error
}

// Find is a method
func (u *PersonalInfoRepo) Find(ctx context.Context, id int) (*entity.PersonalInfo, error) {
	var personalInfo entity.PersonalInfo
	err := u.db.WithContext(ctx).First(&personalInfo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &personalInfo, nil
}

// Get is a method
func (u *PersonalInfoRepo) Get(ctx context.Context) ([]entity.PersonalInfo, error) {
	var personalInfos []entity.PersonalInfo

	err := u.db.WithContext(ctx).Find(&personalInfos).Error
	if err != nil {
		return nil, err
	}

	return personalInfos, nil
}

// Delete is a method
func (u *PersonalInfoRepo) Delete(ctx context.Context, id int) error {
	return u.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.PersonalInfo{}).Error
}

// FindByEmail is a method
func (u *PersonalInfoRepo) FindByEmail(ctx context.Context, blindIndex string) (*entity.PersonalInfo, error) {
	var user entity.PersonalInfo

	err := u.db.WithContext(ctx).First(&user, "email_blind_index = ?", blindIndex).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByPhone is a method
func (u *PersonalInfoRepo) FindByPhone(ctx context.Context, blindIndex string) (*entity.PersonalInfo, error) {
	var personalInfo entity.PersonalInfo

	err := u.db.WithContext(ctx).First(&personalInfo, "phone_blind_index = ?", blindIndex).Error
	if err != nil {
		return nil, err
	}

	return &personalInfo, nil
}
