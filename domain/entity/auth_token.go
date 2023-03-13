package entity

import (
	"time"

	"gorm.io/gorm"
)

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
