package entity

import (
	"github/achjailani/go-simple-grpc/domain/contract"
	"time"

	"gorm.io/gorm"
)

// User is a struct
type User struct {
	ID         uint      `gorm:"not null;uniqueIndex;primaryKey" json:"id"`
	Name       string    `gorm:"size: 100;not null;" json:"name"`
	Username   string    `gorm:"size: 100;null;uniqueIndex" json:"username"`
	Password   string    `gorm:"size: 100;null" json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt
	AuthTokens []AuthToken `gorm:"foreignKey:UserID" json:"user_id"`
}

// TableName is a method
func (u User) TableName() string {
	return "users"
}

// FilterableFields is a method
func (u User) FilterableFields() []interface{} {
	return []interface{}{"id", "name", "username"}
}

// TimeFields is a method
func (u User) TimeFields() []interface{} {
	return []interface{}{"created_at", "updated_at", "deleted_at"}
}

var _ contract.EntityInterface = &User{}
