package storagex_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/storagex"
	"github/achjailani/go-simple-grpc/tests"
	"testing"
)

func TestNewMinio(t *testing.T) {
	textbox := tests.Init()

	r, err := storagex.NewMinio(textbox.Cfg)

	assert.NoError(t, err)
	assert.NotNil(t, r)
}
