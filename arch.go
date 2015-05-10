package goreutils

import (
	"runtime"
)

// Arch prints the machine hardware name.
func Arch() string {
	return runtime.GOARCH
}
