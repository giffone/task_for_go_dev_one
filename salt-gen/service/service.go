package service

import (
	"crypto/rand"
	"fmt"
)

func Generate(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%x", b)

	// return base64.StdEncoding.EncodeToString(b)[:length], nil
	return s[:length], nil
}
