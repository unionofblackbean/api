package pools

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBytesBuffer(t *testing.T) {
	buf := GetBytesBuffer()
	assert.NotNil(t, buf)
	assert.IsType(t, new(bytes.Buffer), buf)
}

func TestPutBytesBuffer(t *testing.T) {
	oriBuf := new(bytes.Buffer)
	PutBytesBuffer(oriBuf)

	buf := GetBytesBuffer()
	assert.Same(t, oriBuf, buf)
}
