package entity

import (
	"github/achjailani/go-simple-grpc/domain/contract"
	"time"

	"gorm.io/gorm"
)

// AuthToken is a struct
type AuthToken struct {
	ID           uint      `gorm:"not null;uniqueIndex;primaryKey"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	AccessToken  string    `gorm:"null" json:"access_token"`
	RefreshToken string    `gorm:"null" json:"refresh_token"`
	ExpiredAt    time.Time `gorm:"not null" json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt
}

// TableName is a method
func (u AuthToken) TableName() string {
	return "auth_tokens"
}

// FilterableFields is a method
func (u AuthToken) FilterableFields() []interface{} {
	return []interface{}{"id", "user_id", "access_token", "refresh_token"}
}

// TimeFields is a method
func (u AuthToken) TimeFields() []interface{} {
	return []interface{}{"expired_at", "created_at", "updated_at", "deleted_at"}
}

var _ contract.EntityInterface = &AuthToken{}
