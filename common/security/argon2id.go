package security

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

type Argon2idParams struct {
	Version     int
	SaltLength  int
	Time        uint32
	Memory      uint32
	Parallelism uint8
	HashLength  uint32
}

var DefaultArgon2idParams = &Argon2idParams{
	Version:     argon2.Version,
	SaltLength:  16,
	Time:        3,
	Memory:      1024 * 64,
	Parallelism: 4,
	HashLength:  32,
}

func Argon2idHashPassword(password []byte, params *Argon2idParams) ([]byte, []byte, error) {
	salt, err := GenerateSalt(params.SaltLength)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate salt -> %v", err)
	}

	pwHash := argon2.IDKey(
		password, salt,
		params.Time, params.Memory, params.Parallelism,
		params.HashLength)

	return salt, pwHash, nil
}

func Argon2idFormatPasswordHash(salt, pwHash []byte, params *Argon2idParams) string {
	return fmt.Sprintf(
		"$argon2i$v=%d$m=%d,t=%d,p=%d$%s$%s",
		params.Version,
		params.Memory,
		params.Time,
		params.Parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(pwHash))
}
