package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

var testConfigString = `{"rest":{"bind_addr":"127.0.0.1","bind_port":8080}}`

func checkConfigValues(t *testing.T, config *Config) {
	assert.Equal(t, "127.0.0.1", config.Rest.BindAddr)
	assert.EqualValues(t, 8080, config.Rest.BindPort)
}

func TestParse(t *testing.T) {
	config, err := Parse([]byte(testConfigString))
	assert.NotNil(t, config)
	assert.Nil(t, err)

	checkConfigValues(t, config)
}

func TestParseString(t *testing.T) {
	config, err := ParseString(testConfigString)
	assert.NotNil(t, config)
	assert.Nil(t, err)

	checkConfigValues(t, config)
}

func TestParseFile(t *testing.T) {
	testConfigFilename := filepath.Join(t.TempDir(), "config.json")
	testConfigFile, err := os.Create(testConfigFilename)
	if err != nil {
		t.Errorf("failed to create test config -> %v", err)
		t.FailNow()
	}

	_, err = testConfigFile.WriteString(testConfigString)
	if err != nil {
		t.Errorf("failed to write test config -> %v", err)
		t.FailNow()
	}

	config, err := ParseFile(testConfigFilename)

	checkConfigValues(t, config)
}
