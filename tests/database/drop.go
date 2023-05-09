package database

import (
	"context"
	"github/achjailani/go-simple-grpc/infrastructure/registry"
	"gorm.io/gorm"
)

// Drop is struct to hold connection to DB
type Drop struct {
	db *gorm.DB
}

// NewDrop is constructor
func NewDrop(db *gorm.DB) *Drop {
	return &Drop{db: db}
}

// DropPostgresql is function drop all tables of postgresql
func (op *Drop) DropPostgresql(ctx context.Context) error {
	for _, table := range registry.CollectTables() {
		ok := op.db.WithContext(ctx).Migrator().HasTable(table.Name)
		if ok {
			err := op.db.WithContext(ctx).Migrator().DropTable(table.Name)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Reset is function drop all tables & recreate them
func (op *Drop) Reset(ctx context.Context) error {
	// disable foreign key constraint
	errDis := op.db.WithContext(context.Background()).Exec("DROP SCHEMA public CASCADE;").Error
	if errDis != nil {
		return errDis
	}

	err := op.DropPostgresql(ctx)
	if err != nil {
		return err
	}

	// enable foreign key constraint
	errEne := op.db.WithContext(context.Background()).Exec("CREATE SCHEMA public;").Error
	if errEne != nil {
		return errEne
	}

	err = registry.NewEntityRegistry().AutoMigrate(op.db)
	if err != nil {
		return err
	}

	return nil
}
