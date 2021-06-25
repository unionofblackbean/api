package jws

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Header map[string]interface{}

func NewHeader() Header {
	return make(Header)
}

func (h Header) Get(key string) interface{} {
	return h[key]
}

func (h Header) Set(key string, value interface{}) {
	h[key] = value
}

func (h Header) Remove(key string) {
	delete(h, key)
}

func (h Header) Exists(key string) bool {
	_, found := h[key]
	return found
}

func (h Header) Encode() ([]byte, error) {
	hJson, err := json.Marshal(h)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal header -> %v", err)
	}

	hEncoded := make([]byte, (len(hJson)*8-1)/6+1)
	base64.RawURLEncoding.Encode(hEncoded, hJson)

	return hEncoded, nil
}
