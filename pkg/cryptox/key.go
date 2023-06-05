package cryptox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

const (
	// rsaKeySize is a constant size
	rsaKeySize = 2048
)

// GenerateKeyPair is a method
func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return nil, nil, err
	}

	return privateKey, &privateKey.PublicKey, nil
}

// GenerateKeyBytes is function to generate bytes key
func GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
	pri, pub, err := GenerateKeyPair()
	if err != nil {
		return nil, nil, err
	}

	priBytes, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		return nil, nil, err
	}

	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, nil, err
	}

	return priBytes, pubBytes, nil
}

// SavePrivateKeyToFile is a function
func SavePrivateKeyToFile(privateKey *rsa.PrivateKey, filename string) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	keyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	return pem.Encode(file, keyPem)
}

// SavePublicKeyToFile is a function
func SavePublicKeyToFile(publicKey *rsa.PublicKey, filename string) error {
	keyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	keyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: keyBytes,
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, keyPem)
}
