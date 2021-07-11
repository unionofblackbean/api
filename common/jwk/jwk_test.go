package jwk

import (
	"github.com/leungyauming/api/common/jwa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	token := New()
	assert.NotNil(t, token)
}

func TestJWK_Get(t *testing.T) {
	token := New()
	token.Set(ParamAlg, jwa.AlgES256)

	assert.Equal(t, jwa.AlgES256, token.Get(ParamAlg))
}

func TestJWK_Set(t *testing.T) {
	token := New()
	token.Set(ParamAlg, jwa.AlgES256)

	assert.True(t, token.Exists(ParamAlg))
	assert.Equal(t, jwa.AlgES256, token.Get(ParamAlg))
}

func TestJWK_Remove(t *testing.T) {
	token := New()
	token.Set(ParamAlg, jwa.AlgES256)
	token.Remove(ParamAlg)

	assert.False(t, token.Exists(ParamAlg))
}

func TestJWK_Exists(t *testing.T) {
	token := New()
	token.Set(ParamAlg, jwa.AlgES256)

	assert.True(t, token.Exists(ParamAlg))
}

func TestJWK_Build(t *testing.T) {
	token := New()
	token.Set(ParamAlg, jwa.AlgES256)

	tj, err := token.Build()
	assert.Nil(t, err)
	assert.Equal(t, `{"alg":"ES256"}`, string(tj))
}
