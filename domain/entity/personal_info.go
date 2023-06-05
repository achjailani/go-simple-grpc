package entity

import (
	"github/achjailani/go-simple-grpc/domain/contract"
	"time"
)

// PersonalInfo is a struct
type PersonalInfo struct {
	ID              uint      `gorm:"not null;uniqueIndex;primaryKey" json:"id"`
	Name            string    `gorm:"column:name;size:100;not null;" json:"name"`
	Email           string    `gorm:"column:email;size:100;not null" json:"email"`
	Phone           string    `gorm:"column:phone;size:100;not null" json:"phone"`
	NameBlindIndex  string    `gorm:"column:name_blind_index;size:150;index" json:"name_blind_index"`
	EmailBlindIndex string    `gorm:"column:email_blind_index;size:150;index" json:"email_blind_index"`
	PhoneBlindIndex string    `gorm:"column:phone_blind_index;size:150;index" json:"phone_blind_index"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TableName is a method
func (u PersonalInfo) TableName() string {
	return "personal_infos"
}

// FilterableFields is a method
func (u PersonalInfo) FilterableFields() []interface{} {
	return []interface{}{"id", "name", "username"}
}

// TimeFields is a method
func (u PersonalInfo) TimeFields() []interface{} {
	return []interface{}{"created_at", "updated_at", "deleted_at"}
}

var _ contract.EntityInterface = &User{}
