package jws

import (
	"crypto/hmac"
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
		return nil, fmt.Errorf("failed to write jws to sign function -> %v", err)
	}
	return signHash.Sum(nil), nil
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
	}

	if err != nil {
		return nil, fmt.Errorf("failed to sign jws -> %v", err)
	}

	return sign, nil
}
