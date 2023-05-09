package entity

import (
	"fmt"
	"github/achjailani/go-simple-grpc/domain/contract"
	"gorm.io/gorm"
	"log"
	"os"
)

// Registry is a struct which hold entities and their tables
type Registry struct {
	Entities []contract.Entity
	Tables   []contract.Table
}

var _ contract.RegistryInterface = &Registry{}

// AutoMigrate is a method to migrate
func (r *Registry) AutoMigrate(db *gorm.DB) error {
	var err error

	for _, model := range r.Entities {
		err = db.AutoMigrate(model.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}

// ResetDatabase is a method to reset database
func (r *Registry) ResetDatabase(db *gorm.DB) error {
	var err error

	if os.Getenv("APP_ENV") == "production" {
		return nil
	}

	for _, table := range r.Tables {
		errDrop := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table.Name))
		if errDrop != nil {
			log.Fatal(errDrop)
		}
	}

	for _, model := range r.Entities {
		err = db.AutoMigrate(model.Entity)
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}
