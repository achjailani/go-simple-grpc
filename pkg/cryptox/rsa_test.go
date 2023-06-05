package cryptox_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/cryptox"
	"testing"
)

func TestRSA(t *testing.T) {
	privateKey, publicKey, err := cryptox.GenerateKeyPair()

	assert.NoError(t, err)
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)

	data := []byte("Hi, this message should be encrypted, because it contains all about us, no one should know it :)")

	encrypted, err := cryptox.RSAEncrypt(publicKey, data)
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)

	decrypted, err := cryptox.RSADecrypt(privateKey, encrypted)
	assert.NoError(t, err)
	assert.NotEmpty(t, decrypted)
}
