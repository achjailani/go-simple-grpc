package entity

import (
	"gorm.io/gorm"
	"time"
)

// HttpLog is a struct
type HttpLog struct {
	ID        uint           `gorm:"column:id;not null;uniqueIndex;primaryKey" json:"id"`
	Ip        string         `gorm:"column:ip;size:100"`
	Path      string         `gorm:"column:path;size:200"`
	Method    string         `gorm:"column:method"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
