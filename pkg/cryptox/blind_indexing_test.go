package cryptox_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/cryptox"
	"testing"
)

func TestMakeBlindIndex(t *testing.T) {
	data := "087750676800"
	r, err := cryptox.MakeBlindIndex(data)

	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	fmt.Println(r)

	data2 := "Hello world, today is a great day for coding"

	r, _ = cryptox.MakeBlindIndex(data2)
	fmt.Println(r)
}
