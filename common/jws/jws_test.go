package jws

import (
	"github.com/leungyauming/api/common/jwa"
	"github.com/leungyauming/api/common/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	jws := New()
	assert.NotNil(t, jws)
	assert.NotNil(t, jws.Header)
	assert.NotNil(t, jws.Payload)
}

func TestJWS_Sign_HS256(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS256)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign("thisissecret")
	assert.Nil(t, err)
	assert.Equal(t, []byte("300qEKK0Iibf-_t8rXzVDdBPX-xEXT-9v_sFJ2X4zYU"), sign)
}

func TestJWS_Sign_HS384(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS384)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign("thisissecret")
	assert.Nil(t, err)
	assert.Equal(t, []byte("z5LD-Pb68M5VN3y61AzPiRN_pa-N2Vcz6i1wONCulZo5AsrATUMGxx4_OKxh1fGo"), sign)
}

func TestJWS_Sign_HS512(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS512)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign("thisissecret")
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte("d427lJYZVKAxVqnmMEXsWRLWzGUeAVgIWKggs0sbI1le40QbI4e0e3eAYrKwQ0OcBb2DbS6weRllMfYKaPsAHg"),
		sign)
}
