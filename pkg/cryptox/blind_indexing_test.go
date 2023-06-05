package cryptox_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/cryptox"
	"testing"
)

func TestMakeBlindIndex(t *testing.T) {
	data := "087750676800"
	r, err := cryptox.MakeBlindIndex(data)

	assert.NoError(t, err)
	assert.NotEmpty(t, r)
}
