package security

import (
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
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

func Argon2idHashPassword(password string, params *Argon2idParams) ([]byte, []byte, error) {
	salt, err := GenerateSalt(params.SaltLength)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate salt -> %v", err)
	}

	return salt, argon2.IDKey(
		[]byte(password), salt,
		params.Time, params.Memory, params.Parallelism,
		params.HashLength), nil
}

func Argon2idHashPasswordWithSalt(salt []byte, password string, params *Argon2idParams) ([]byte, error) {
	if len(salt) != params.SaltLength {
		return nil, errors.New("mismatched salt length")
	}

	return argon2.IDKey(
		[]byte(password), salt,
		params.Time, params.Memory, params.Parallelism,
		params.HashLength), nil
}

func Argon2idVerifyPassword(password string, encodedHash string) (bool, error) {
	salt, _, params, err := Argon2idDecodePasswordHash(encodedHash)
	if err != nil {
		return false, fmt.Errorf("failed to decode password hash -> %v", err)
	}

	pwHash, err := Argon2idHashPasswordWithSalt(salt, password, params)
	if err != nil {
		return false, fmt.Errorf("failed to hash password with salt -> %v", err)
	}
	pwHashEncoded := Argon2idEncodePasswordHash(salt, pwHash, params)

	return pwHashEncoded == encodedHash, nil
}

func Argon2idEncodePasswordHash(salt, pwHash []byte, params *Argon2idParams) string {
	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		params.Version,
		params.Memory,
		params.Time,
		params.Parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(pwHash))
}

func Argon2idDecodePasswordHash(encoded string) ([]byte, []byte, *Argon2idParams, error) {
	encodedSlices := strings.Split(encoded, "$")

	if encodedSlices[1] != "argon2id" {
		return nil, nil, nil, errors.New("argon2 type is not argon2id")
	}

	//$TYPE$v=VERSION$m=MEMORY,t=TIME,p=PARALLELISM$SALT$HASH
	var (
		version       int
		memory        uint32
		time          uint32
		parallelism   uint8
		saltEncoded   string
		pwHashEncoded string
	)

	n, err := fmt.Sscanf(encodedSlices[2], "v=%d", &version)
	if n != 1 || err != nil {
		return nil, nil, nil, errors.New("failed to scan version")
	}

	n, err = fmt.Sscanf(encodedSlices[3], "m=%d,t=%d,p=%d", &memory, &time, &parallelism)
	if n != 3 || err != nil {
		return nil, nil, nil, errors.New("failed to scan memory, time and parallelism")
	}

	saltEncoded = encodedSlices[4]
	pwHashEncoded = encodedSlices[5]

	salt, err := base64.RawStdEncoding.DecodeString(saltEncoded)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode salt -> %v", err)
	}
	pwHash, err := base64.RawStdEncoding.DecodeString(pwHashEncoded)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode hash -> %v", err)
	}

	return salt, pwHash,
		&Argon2idParams{
			Version:     version,
			SaltLength:  len(salt),
			Time:        time,
			Memory:      memory,
			Parallelism: parallelism,
			HashLength:  uint32(len(pwHash)),
		}, nil
}
