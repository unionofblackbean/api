package jws

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/leungyauming/api/common/jwa"
	"github.com/leungyauming/api/common/pools"
	"hash"
)

type JWS struct {
	Header  Header
	Payload Payload
}

func New() *JWS {
	return &JWS{
		Header:  NewHeader(),
		Payload: NewPayload(),
	}
}

func (jws *JWS) encodeHeaderPayload() ([]byte, error) {
	headerEncoded, err := jws.Header.Encode()
	if err != nil {
		return nil, fmt.Errorf("failed to encode header -> %v", err)
	}

	payloadEncoded, err := jws.Payload.Encode()
	if err != nil {
		return nil, fmt.Errorf("failed to encode payload -> %v", err)
	}

	buf := pools.GetBytesBuffer()
	buf.Write(headerEncoded)
	buf.WriteRune('.')
	buf.Write(payloadEncoded)

	headerPayloadEncoded := buf.Bytes()
	pools.PutBytesBuffer(buf)

	return headerPayloadEncoded, nil
}

func (jws *JWS) signHS(iSecret interface{}) ([]byte, error) {
	var secret []byte
	switch iSecret.(type) {
	case string:
		secret = []byte(iSecret.(string))
	case []byte:
		secret = iSecret.([]byte)
	default:
		return nil, errors.New("unknown secret data type")
	}

	var hashFunc func() hash.Hash
	switch jws.Header.Get("alg") {
	case jwa.AlgHS256:
		hashFunc = sha256.New
	case jwa.AlgHS384:
		hashFunc = sha512.New384
	case jwa.AlgHS512:
		hashFunc = sha512.New
	}

	headerPayloadEncoded, err := jws.encodeHeaderPayload()
	if err != nil {
		return nil, fmt.Errorf("failed to encode header and payload -> %v", err)
	}

	signHash := hmac.New(hashFunc, secret)
	_, err = signHash.Write(headerPayloadEncoded)
	if err != nil {
		return nil, fmt.Errorf("failed to write header and payload to HMAC sign function -> %v", err)
	}
	return signHash.Sum(nil), nil
}

func (jws *JWS) signRS(iPriKey interface{}) ([]byte, error) {
	priKey, ok := iPriKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("unknown private key data type -> %v")
	}

	headerPayloadEncoded, err := jws.encodeHeaderPayload()
	if err != nil {
		return nil, fmt.Errorf("failed to encode header and payload -> %v", err)
	}

	var hashFunc crypto.Hash
	var hashed []byte
	switch jws.Header.Get("alg") {
	case jwa.AlgHS256:
		hashFunc = crypto.SHA256
		tmpHashed := sha256.Sum256(headerPayloadEncoded)
		hashed = tmpHashed[:]
	case jwa.AlgHS384:
		hashFunc = crypto.SHA384
		tmpHashed := sha512.Sum384(headerPayloadEncoded)
		hashed = tmpHashed[:]
	case jwa.AlgHS512:
		hashFunc = crypto.SHA512
		tmpHashed := sha512.Sum512(headerPayloadEncoded)
		hashed = tmpHashed[:]
	}

	sign, err := rsa.SignPKCS1v15(rand.Reader, priKey, hashFunc, hashed)
	if err != nil {
		return nil, fmt.Errorf("failed to sign header and payload using RSA PKCS #1 v1.5 -> %v", err)
	}

	return sign, nil
}

func (jws *JWS) signES(iPriKey interface{}) ([]byte, error) {
	priKey, ok := iPriKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("unknown private key data type")
	}

	headerPayloadEncoded, err := jws.encodeHeaderPayload()
	if err != nil {
		return nil, fmt.Errorf("failed to encode header and payload -> %v", err)
	}

	var hashed []byte
	switch jws.Header.Get("alg") {
	case jwa.AlgHS256:
		tmpHashed := sha256.Sum256(headerPayloadEncoded)
		hashed = tmpHashed[:]
	case jwa.AlgHS384:
		tmpHashed := sha512.Sum384(headerPayloadEncoded)
		hashed = tmpHashed[:]
	case jwa.AlgHS512:
		tmpHashed := sha512.Sum512(headerPayloadEncoded)
		hashed = tmpHashed[:]
	}

	sign, err := ecdsa.SignASN1(rand.Reader, priKey, hashed)
	if err != nil {
		return nil, fmt.Errorf("failed to sign header and payload using ECDSA -> %v", err)
	}

	return sign, nil
}

func (jws *JWS) signPS(iPriKey interface{}) ([]byte, error) {
	priKey, ok := iPriKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("unknown private key data type")
	}

	headerPayloadEncoded, err := jws.encodeHeaderPayload()
	if err != nil {
		return nil, fmt.Errorf("failed to encode header and payload -> %v", err)
	}

	var hashFunc crypto.Hash
	var hashed []byte
	switch jws.Header.Get("alg") {
	case jwa.AlgHS256:
		hashFunc = crypto.SHA256
		tmpHashed := sha256.Sum256(headerPayloadEncoded)
		hashed = tmpHashed[:]
	case jwa.AlgHS384:
		hashFunc = crypto.SHA384
		tmpHashed := sha512.Sum384(headerPayloadEncoded)
		hashed = tmpHashed[:]
	case jwa.AlgHS512:
		hashFunc = crypto.SHA512
		tmpHashed := sha512.Sum512(headerPayloadEncoded)
		hashed = tmpHashed[:]
	}

	sign, err := rsa.SignPSS(rand.Reader, priKey, hashFunc, hashed, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to sign header and payload with RSA using PSS -> %v", err)
	}

	return sign, nil
}

func (jws *JWS) Sign(secretOrPriKey interface{}) ([]byte, error) {
	if !jws.Header.Exists(HeaderParamAlg) {
		return nil, errors.New("missing alg header parameter")
	}

	var sign []byte
	var err error

	alg := jws.Header.Get(HeaderParamAlg)
	switch alg {
	case jwa.AlgHS256, jwa.AlgHS384, jwa.AlgHS512:
		sign, err = jws.signHS(secretOrPriKey)
	case jwa.AlgRS256, jwa.AlgRS384, jwa.AlgRS512:
		sign, err = jws.signRS(secretOrPriKey)
	case jwa.AlgES256, jwa.AlgES384, jwa.AlgES512:
		sign, err = jws.signES(secretOrPriKey)
	case jwa.AlgPS256, jwa.AlgPS384, jwa.AlgPS512:
		sign, err = jws.signPS(secretOrPriKey)
	default:
		return nil, errors.New("unknown alg header parameter option")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to sign jws -> %v", err)
	}

	return sign, nil
}
