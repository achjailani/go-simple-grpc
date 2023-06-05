package cryptox

import (
	"crypto/rand"
	"crypto/rsa"
)

// RSAEncrypt is a function
func RSAEncrypt(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

// RSADecrypt is a function
func RSADecrypt(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
}
