package jwk

import (
	"github.com/stretchr/testify/assert"
	"github.com/unionofblackbean/api/common/jwa"
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet()
	assert.NotNil(t, set)
}

func TestSet_AddKey(t *testing.T) {
	jwk := New()
	jwk.Set(ParamAlg, jwa.AlgES256)

	set := NewSet()
	set.AddKey(jwk)
	assert.ElementsMatch(t, []JWK{jwk}, set.Keys)
}

func TestSet_RemoveKey(t *testing.T) {
	jwk := New()
	jwk.Set(ParamAlg, jwa.AlgES256)

	set := NewSet()
	set.AddKey(jwk)
	set.RemoveKey(0)
	assert.Empty(t, set.Keys)
}

func TestSet_Build(t *testing.T) {
	jwk := New()
	jwk.Set(ParamAlg, jwa.AlgES256)

	set := NewSet()
	set.AddKey(jwk)

	sj, err := set.Build()
	assert.Nil(t, err)
	assert.Equal(t, `{"keys":[{"alg":"ES256"}]}`, string(sj))
}
