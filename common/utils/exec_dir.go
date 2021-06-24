package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExecutableDirectory() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path -> %v", err)
	}

	path, err := filepath.EvalSymlinks(execPath)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate possible symlink of executable -> %v", err)
	}

	return filepath.Dir(path), nil
}
