package config

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var testConfigStringReader = strings.NewReader(testConfigString)

func TestNewParser(t *testing.T) {
	parser := NewParser(testConfigStringReader)
	assert.NotNil(t, parser.dec)
}

func TestParser_Parse(t *testing.T) {
	parser := NewParser(testConfigStringReader)
	config, err := parser.Parse()
	assert.NotNil(t, config)
	assert.Nil(t, err)

	checkConfigValues(t, config)
}
