package security

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomBytes(len int) ([]byte, error) {
	salt := make([]byte, len)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to read random bytes from generator -> %v", err)
	}

	return salt, nil
}
