package jwk

import (
	"encoding/json"
	"fmt"
)

type Set struct {
	Keys []JWK `json:"keys"`
}

func (s *Set) AddKey(jwk JWK) {
	s.Keys = append(s.Keys, jwk)
}

func (s *Set) RemoveKey(index int) {
	s.Keys = append(s.Keys[:index], s.Keys[index+1:]...)
}

func (s *Set) Build() ([]byte, error) {
	sj, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal jwk set -> %v", err)
	}

	return sj, err
}

func NewSet() *Set {
	return new(Set)
}
