package entity

import (
	"github/achjailani/go-simple-grpc/domain/contract"
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

// TableName is a method
func (u HttpLog) TableName() string {
	return "http_logs"
}

// FilterableFields is a method
func (u HttpLog) FilterableFields() []interface{} {
	return []interface{}{"id", "ip", "path", "method"}
}

// TimeFields is a method
func (u HttpLog) TimeFields() []interface{} {
	return []interface{}{"created_at", "updated_at", "deleted_at"}
}

var _ contract.EntityInterface = &HttpLog{}
