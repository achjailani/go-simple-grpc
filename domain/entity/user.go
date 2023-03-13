package entity

import (
	"time"

	"gorm.io/gorm"
)

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
