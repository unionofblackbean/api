package jws

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader_Get(t *testing.T) {
	h := NewHeader()

	assert.Nil(t, h.Get("alg"))
	h.Set("alg", "RS256")
	assert.NotNil(t, h.Get("alg"))
}

func TestHeader_Set(t *testing.T) {
	h := NewHeader()

	h.Set("alg", "RS256")
	assert.True(t, h.Exists("alg"))
}

func TestHeader_Remove(t *testing.T) {
	h := NewHeader()

	h.Set("alg", "RS256")
	assert.True(t, h.Exists("alg"))

	h.Remove("alg")
	assert.False(t, h.Exists("alg"))
}

func TestHeader_Exists(t *testing.T) {
	h := NewHeader()

	assert.False(t, h.Exists("alg"))
	h.Set("alg", "RS256")
	assert.True(t, h.Exists("alg"))
}
