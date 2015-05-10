package goreutils

import (
	"path/filepath"
	"strings"
)

// Basename strips directory and suffix from filenames.
func Basename(path, suffix string) string {
	base := filepath.Base(path)

	if strings.Contains(suffix, "/") || base == suffix {
		return base
	}

	return strings.TrimSuffix(base, suffix)
}
