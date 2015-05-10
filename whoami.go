package goreutils

import (
	"os/user"
)

// Whoami returns the effective userid.
func Whoami() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}

	return u.Username, nil
}
