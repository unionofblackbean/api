package utils

import (
	"os"
)

func IsFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
