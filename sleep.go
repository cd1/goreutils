package goreutils

import (
	"time"
)

// Sleep delays for a specified amount of time.
func Sleep(duration time.Duration) {
	time.Sleep(duration)
}
