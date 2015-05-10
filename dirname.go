package goreutils

import (
	"path/filepath"
)

// Dirname strips last component from file name.
func Dirname(path string) string {
	return filepath.Dir(path)
}
