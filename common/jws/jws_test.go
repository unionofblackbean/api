package jws

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	jws := New()
	assert.NotNil(t, jws)
	assert.NotNil(t, jws.Header)
	assert.NotNil(t, jws.Payload)
}
