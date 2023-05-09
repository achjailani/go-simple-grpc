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

func NewDBConnection(cfg *config.Config) (*gorm.DB, error) {
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

	if cfg.DBConfig.DBLog {
		gormConfig.Logger = newLogger
	}

	var dbURL string
	var db *gorm.DB

	switch cfg.DBConfig.DBDriver {
	case driverPostgres:
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			cfg.DBConfig.DBHost,
			cfg.DBConfig.DBUser,
			cfg.DBConfig.DBPassword,
			cfg.DBConfig.DBName,
			cfg.DBConfig.DBPort,
			cfg.DBConfig.DBTimeZone,
		)

		if cfg.TestMode {
			dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
				cfg.DBTestConfig.DBHost,
				cfg.DBTestConfig.DBUser,
				cfg.DBTestConfig.DBPassword,
				cfg.DBTestConfig.DBName,
				cfg.DBTestConfig.DBPort,
				cfg.DBTestConfig.DBTimeZone,
			)
		}

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
		&entity.User{},
		&entity.AuthToken{},
		&entity.HttpLog{},
	)

	if err != nil {
		return err
	}

	return nil
}
