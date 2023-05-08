package persistence

import (
	"fmt"
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/entity"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	driverPostgres = "postgres"
)

func NewDBConnection(config config.DBConfig) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if config.DBLog {
		gormConfig.Logger = newLogger
	}

	var dbURL string
	var db *gorm.DB

	switch config.DBDriver {
	case driverPostgres:
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			config.DBHost,
			config.DBUser,
			config.DBPassword,
			config.DBName,
			config.DBPort,
			config.DBTimeZone,
		)
		dbConn, err := gorm.Open(postgres.Open(dbURL), gormConfig)
		if err != nil {
			return nil, err
		}

		return dbConn, nil
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.AuthToken{},
		&entity.User{},
		&entity.HttpLog{},
	)

	if err != nil {
		return err
	}

	return nil
}
