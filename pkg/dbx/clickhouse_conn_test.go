package dbx_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/dbx"
	"github/achjailani/go-simple-grpc/tests"
	"testing"
)

func TestNewClickhouseConn(t *testing.T) {
	testbox := tests.Init()
	conn, err := dbx.NewClickhouseConn(testbox.Cfg)

	assert.NoError(t, err)
	assert.NotNil(t, conn)
	v, err := conn.ServerVersion()
	assert.NoError(t, err)
	fmt.Println(v.String())
}
