package cryptox

import (
	"crypto/sha256"
	"encoding/hex"
)

// MakeBlindIndex is a function
func MakeBlindIndex(data string) (string, error) {
	hash := sha256.Sum256([]byte(data))
	// hash[:] convert array to slice
	blindIndex := hex.EncodeToString(hash[:])

	return blindIndex, nil
}
