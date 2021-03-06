package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestIsFileExists(t *testing.T) {
	tempDir := t.TempDir()

	testFilename := filepath.Join(tempDir, "test.txt")
	testFile, err := os.Create(testFilename)
	if err != nil {
		t.Errorf("failed to create test file -> %v", err)
	}
	defer testFile.Close()

	assert.True(t, IsFileExists(testFilename))

	randomFilename := filepath.Join(tempDir, "random.txt")
	assert.False(t, IsFileExists(randomFilename))
}
