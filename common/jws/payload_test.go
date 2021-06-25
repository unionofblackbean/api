package jws

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPayload(t *testing.T) {
	p := NewPayload()
	assert.NotNil(t, p)
}

func TestPayload_Get(t *testing.T) {
	p := NewPayload()

	assert.Nil(t, p.Get("role"))
	p.Set("role", "admin")
	assert.NotNil(t, p.Get("role"))
}

func TestPayload_Set(t *testing.T) {
	p := NewPayload()

	p.Set("role", "admin")
	assert.True(t, p.Exists("role"))
}

func TestPayload_Remove(t *testing.T) {
	p := NewPayload()

	p.Set("role", "admin")
	assert.True(t, p.Exists("role"))

	p.Remove("role")
	assert.False(t, p.Exists("role"))
}

func TestPayload_Exists(t *testing.T) {
	p := NewPayload()

	assert.False(t, p.Exists("role"))
	p.Set("role", "admin")
	assert.True(t, p.Exists("role"))
}

func TestPayload_Encode(t *testing.T) {
	p := NewPayload()
	p.Set("role", "admin")

	pEncoded, err := p.Encode()
	assert.Nil(t, err)

	assert.Equal(t, []byte("eyJyb2xlIjoiYWRtaW4ifQ"), pEncoded)
}
