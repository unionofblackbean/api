package jwk

import (
	"encoding/json"
	"fmt"
)

type JWK map[string]interface{}

func New() JWK {
	return make(JWK)
}

func (jwk JWK) Get(key string) interface{} {
	return jwk[key]
}

func (jwk JWK) Set(key string, value interface{}) {
	jwk[key] = value
}

func (jwk JWK) Remove(key string) {
	delete(jwk, key)
}

func (jwk JWK) Exists(key string) bool {
	_, found := jwk[key]
	return found
}

func (jwk JWK) Build() ([]byte, error) {
	kj, err := json.Marshal(jwk)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal jwk -> %v", err)
	}

	return kj, nil
}
