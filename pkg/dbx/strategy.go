package dbx

import (
	"context"
	"database/sql"
)

type Adapter interface {
	Ping() error
	InTransaction() bool
	Close() error
	Query(ctx context.Context, dst any, query string, args ...any) error
	QueryRow(ctx context.Context, dst any, query string, args ...any) error
	QueryX(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowX(ctx context.Context, query string, args ...any) *sql.Row
	Exec(ctx context.Context, query string, args ...any) (_ int64, err error)
	//Transact(ctx context.Context, iso sql.IsolationLevel, txFunc func(*DB) error) (err error)
}
