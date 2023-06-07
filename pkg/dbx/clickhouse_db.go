package dbx

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github/achjailani/go-simple-grpc/config"
)

// ClickhouseSQL is a struct
type ClickhouseSQL struct {
	sql driver.Conn
}

// NewClickhouseSQL is a constructor
func NewClickhouseSQL(cfg *config.Config) (*ClickhouseSQL, error) {
	conn, err := NewClickhouseConn(cfg)
	if err != nil {
		return nil, err
	}

	return &ClickhouseSQL{
		sql: conn,
	}, nil
}
