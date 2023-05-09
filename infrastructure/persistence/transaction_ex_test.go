package persistence_test

import (
	"database/sql"
	"github/achjailani/go-simple-grpc/tests"
	"gorm.io/gorm"
	"testing"
)

func TestTransaction(t *testing.T) {
	box := tests.Init()

	db := box.Dependency.Repo.DB
	tx := db.Begin(&sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	_ = tx
}

func execTransaction(db *gorm.DB) error {
	//TODO implement me
	panic("implement me")
}
