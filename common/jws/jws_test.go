package jws

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/stretchr/testify/assert"
	"github.com/unionofblackbean/api/common/jwa"
	"github.com/unionofblackbean/api/common/jwt"
	"testing"
)

func TestNew(t *testing.T) {
	jws := New()
	assert.NotNil(t, jws)
	assert.NotNil(t, jws.Header)
	assert.NotNil(t, jws.Payload)
}

var hsSecretString = "thisissecret"

func TestJWS_Sign_HS256_1(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS256)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign(hsSecretString)
	assert.Nil(t, err)
	assert.Equal(t, []byte("300qEKK0Iibf-_t8rXzVDdBPX-xEXT-9v_sFJ2X4zYU"), sign)
}

func TestJWS_Sign_HS256_2(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS256)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign([]byte(hsSecretString))
	assert.Nil(t, err)
	assert.Equal(t, []byte("300qEKK0Iibf-_t8rXzVDdBPX-xEXT-9v_sFJ2X4zYU"), sign)
}

func TestJWS_Sign_HS384_1(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS384)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign(hsSecretString)
	assert.Nil(t, err)
	assert.Equal(t, []byte("z5LD-Pb68M5VN3y61AzPiRN_pa-N2Vcz6i1wONCulZo5AsrATUMGxx4_OKxh1fGo"), sign)
}

func TestJWS_Sign_HS384_2(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS384)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign([]byte(hsSecretString))
	assert.Nil(t, err)
	assert.Equal(t, []byte("z5LD-Pb68M5VN3y61AzPiRN_pa-N2Vcz6i1wONCulZo5AsrATUMGxx4_OKxh1fGo"), sign)
}

func TestJWS_Sign_HS512_1(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS512)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign(hsSecretString)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte("d427lJYZVKAxVqnmMEXsWRLWzGUeAVgIWKggs0sbI1le40QbI4e0e3eAYrKwQ0OcBb2DbS6weRllMfYKaPsAHg"),
		sign)
}
func TestJWS_Sign_HS512_2(t *testing.T) {
	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgHS512)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign([]byte(hsSecretString))
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte("d427lJYZVKAxVqnmMEXsWRLWzGUeAVgIWKggs0sbI1le40QbI4e0e3eAYrKwQ0OcBb2DbS6weRllMfYKaPsAHg"),
		sign)
}

var rsaPemPriKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAnzyis1ZjfNB0bBgKFMSvvkTtwlvBsaJq7S5wA+kzeVOVpVWw
kWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHcaT92whREFpLv9cj5lTeJSibyr/Mr
m/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIytvHWTxZYEcXLgAXFuUuaS3uF9gEi
NQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0e+lf4s4OxQawWD79J9/5d3Ry0vbV
3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWbV6L11BWkpzGXSW4Hv43qa+GSYOD2
QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9MwIDAQABAoIBACiARq2wkltjtcjs
kFvZ7w1JAORHbEufEO1Eu27zOIlqbgyAcAl7q+/1bip4Z/x1IVES84/yTaM8p0go
amMhvgry/mS8vNi1BN2SAZEnb/7xSxbflb70bX9RHLJqKnp5GZe2jexw+wyXlwaM
+bclUCrh9e1ltH7IvUrRrQnFJfh+is1fRon9Co9Li0GwoN0x0byrrngU8Ak3Y6D9
D8GjQA4Elm94ST3izJv8iCOLSDBmzsPsXfcCUZfmTfZ5DbUDMbMxRnSo3nQeoKGC
0Lj9FkWcfmLcpGlSXTO+Ww1L7EGq+PT3NtRae1FZPwjddQ1/4V905kyQFLamAA5Y
lSpE2wkCgYEAy1OPLQcZt4NQnQzPz2SBJqQN2P5u3vXl+zNVKP8w4eBv0vWuJJF+
hkGNnSxXQrTkvDOIUddSKOzHHgSg4nY6K02ecyT0PPm/UZvtRpWrnBjcEVtHEJNp
bU9pLD5iZ0J9sbzPU/LxPmuAP2Bs8JmTn6aFRspFrP7W0s1Nmk2jsm0CgYEAyH0X
+jpoqxj4efZfkUrg5GbSEhf+dZglf0tTOA5bVg8IYwtmNk/pniLG/zI7c+GlTc9B
BwfMr59EzBq/eFMI7+LgXaVUsM/sS4Ry+yeK6SJx/otIMWtDfqxsLD8CPMCRvecC
2Pip4uSgrl0MOebl9XKp57GoaUWRWRHqwV4Y6h8CgYAZhI4mh4qZtnhKjY4TKDjx
QYufXSdLAi9v3FxmvchDwOgn4L+PRVdMwDNms2bsL0m5uPn104EzM6w1vzz1zwKz
5pTpPI0OjgWN13Tq8+PKvm/4Ga2MjgOgPWQkslulO/oMcXbPwWC3hcRdr9tcQtn9
Imf9n2spL/6EDFId+Hp/7QKBgAqlWdiXsWckdE1Fn91/NGHsc8syKvjjk1onDcw0
NvVi5vcba9oGdElJX3e9mxqUKMrw7msJJv1MX8LWyMQC5L6YNYHDfbPF1q5L4i8j
8mRex97UVokJQRRA452V2vCO6S5ETgpnad36de3MUxHgCOX3qL382Qx9/THVmbma
3YfRAoGAUxL/Eu5yvMK8SAt/dJK6FedngcM3JEFNplmtLYVLWhkIlNRGDwkg3I5K
y18Ae9n7dHVueyslrb6weq7dTkYDi3iOYRW8HRkIQh06wEdbxt0shTzAJvvCQfrB
jg/3747WSsf/zBTcHihTRBdAv6OmdhV4/dD5YBfLAkLrd+mX7iE=
-----END RSA PRIVATE KEY-----`)

func parseRsaPemPriKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(rsaPemPriKey)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func TestJWS_Sign_RS256(t *testing.T) {
	priKey, err := parseRsaPemPriKey()
	assert.Nil(t, err)

	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgRS256)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign(priKey)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte("YlfN1F8tV9qWiu5ct-xotqV6NRql18wMQZ-yPiZ1JKV0KBCa0JfJ6hXXXzr_qC6bmXi3IupgsxSO99A1HRY4q4M5umA7dti2sVbnYKvdIzmlBt3pvSMu6RRmULdSyxrTNWSxD9s-gte6dQzUPAKSvgBsOd_RVdUJxEtiN0bKzqnWhdxp7yX0DDUTrKCkXjVXYeiRd-PRH6Yn6q_rzZSWXggYs03WmlEa3dinu7G0Wym7WPxiqncAu-FowKf4ESV3AEQMiWl2f-P92t9vLgLXd6kcSiv2QiduPf558J4kw42tEZ_lNdv7ciLwYqHsWC89dDWTRLp7rD8QzlR2DYyOLQ"),
		sign)
}

func TestJWS_Sign_RS384(t *testing.T) {
	priKey, err := parseRsaPemPriKey()
	assert.Nil(t, err)

	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgRS384)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign(priKey)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte("GH4x0nBF2ecH4UJgmInXv7c9qvHtihQNVY57pmdeYIElpidGDHyEOCQCTFCAq9IPTEA0qQa5zHg8R2CkgOaHzgL3HbPqQ759CaiAv1ODv2ZC8_DS814U90n18jJPDbJsg3akUapUV89RQXaWbWPV7_IHfWweDETsha4UDO8gL5CXOTux8SHiX7z5c4ug6O8OuAi0qvCSM7KDp3G6rqSL7ia3w8ZXijDmFr10IDG9-tsCyKcA7gQ7OjQFGfGdnRF7z0CB8Qb5depCEqOpUKlbp7YLa458DlhC0soFh_Dofw6aak6FgiMaUQcPab3s69GyLQLUGfdHE19VSZaBv4BPPA"),
		sign)
}

func TestJWS_Sign_RS512(t *testing.T) {
	priKey, err := parseRsaPemPriKey()
	assert.Nil(t, err)

	jws := New()
	jws.Header.Set(HeaderParamAlg, jwa.AlgRS512)
	jws.Header.Set(HeaderParamTyp, jwt.TypJWT)

	jws.Payload.Set("role", "admin")

	sign, err := jws.Sign(priKey)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte("iK8bfFPB60_cY9oHy9GQOUA6y-wtuq-OM60KBe6Q8HXI_G1gTOXgVDEppM2gCLXO5VM3Wfa7FECK3ZRToc_oKjPhpJweVaubK0J9BtnOwG0GH8uKXYJc4H9Fdv6bHfd42nXeCw8suO-p07SQsz1wRVuWM_xF8z5mjkH_ROCWVciJjDfTfqw3aoK3YvoRFSvRRb5NYL6zvHDqM89kLPOwVeDW_hx3LPbsHvWwvUcKwulXa1zQjobGWbmv2czA5fhvJhdp80kiQyBdRM2m7Z5R2UtpxZmtqbmCh9aPmx0uxnZqiJ_m6NQ7FfeVe6-GX5NOreMU6drgNe3dgafFSjVeuQ"),
		sign)
}
