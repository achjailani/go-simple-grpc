package cryptox_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/cryptox"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	privateKey, publicKey, err := cryptox.GenerateKeyPair()

	assert.NoError(t, err)
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)
}

func TestGenerateKeyBytes(t *testing.T) {
	privateKey, publicKey, err := cryptox.GenerateKeyBytes()

	assert.NoError(t, err)
	assert.NotEmpty(t, privateKey)
	assert.NotEmpty(t, publicKey)
}

func TestGenerateKeyPairSaveToFile(t *testing.T) {
	privateKey, publicKey, err := cryptox.GenerateKeyPair()

	assert.NoError(t, err)
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)

	err = cryptox.SavePrivateKeyToFile(privateKey, "private.key")
	assert.NoError(t, err)
	err = cryptox.SavePublicKeyToFile(publicKey, "public.key")
	assert.NoError(t, err)
}
