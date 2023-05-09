package contract

import "gorm.io/gorm"

// Entity is a struct which hold any entity
type Entity struct {
	Entity interface{}
}

// Table is a struct which hold entity's name
type Table struct {
	Name interface{}
}

// RegistryInterface is a contract
type RegistryInterface interface {
	AutoMigrate(db *gorm.DB) error
	ResetDatabase(db *gorm.DB) error
}
