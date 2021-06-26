package jws

import (
	"fmt"
	"github.com/leungyauming/api/common/pools"
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
