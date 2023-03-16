package client_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/grpc/client"
	"testing"
)

func TestRun(t *testing.T) {
	err := client.Run()

	assert.NoError(t, err)
}
