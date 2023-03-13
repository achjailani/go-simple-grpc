package utils

import (
	"encoding/base64"
	"strings"
)

type Token string

func EncodeBasicAuth(username, password string) string {
	token := base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{username, password}, ":")))

	return token
}

func DecodeBasicAuth(token string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
