package jws

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Payload map[string]interface{}

func NewPayload() Payload {
	return make(Payload)
}

func (p Payload) Get(key string) interface{} {
	return p[key]
}

func (p Payload) Set(key string, value string) {
	p[key] = value
}

func (p Payload) Remove(key string) {
	delete(p, key)
}

func (p Payload) Exists(key string) bool {
	_, found := p[key]
	return found
}

func (p Payload) Encode() ([]byte, error) {
	pJson, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal header -> %v", err)
	}

	pEncoded := make([]byte, (len(pJson)*8-1)/6+1)
	base64.RawURLEncoding.Encode(pEncoded, pJson)

	return pEncoded, nil
}
