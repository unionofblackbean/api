package jws

import (
	"github.com/leungyauming/api/common/jwa"
	"github.com/leungyauming/api/common/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader_Get(t *testing.T) {
	h := NewHeader()

	assert.Nil(t, h.Get(HeaderParamAlg))
	h.Set(HeaderParamAlg, jwa.AlgRS256)
	assert.NotNil(t, h.Get(HeaderParamAlg))
}

func TestHeader_Set(t *testing.T) {
	h := NewHeader()

	h.Set(HeaderParamAlg, jwa.AlgRS256)
	assert.True(t, h.Exists(HeaderParamAlg))
}

func TestHeader_Remove(t *testing.T) {
	h := NewHeader()

	h.Set(HeaderParamAlg, jwa.AlgRS256)
	assert.True(t, h.Exists(HeaderParamAlg))

	h.Remove(HeaderParamAlg)
	assert.False(t, h.Exists(HeaderParamAlg))
}

func TestHeader_Exists(t *testing.T) {
	h := NewHeader()

	assert.False(t, h.Exists(HeaderParamAlg))
	h.Set(HeaderParamAlg, jwa.AlgRS256)
	assert.True(t, h.Exists(HeaderParamAlg))
}

func TestHeader_Encode(t *testing.T) {
	h := NewHeader()

	h.Set(HeaderParamAlg, jwa.AlgHS256)
	h.Set(HeaderParamTyp, jwt.TypJWT)

	hEncoded, err := h.Encode()
	assert.Nil(t, err)

	assert.Equal(t, []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"), hEncoded)
}
