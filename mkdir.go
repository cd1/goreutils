package goreutils

import (
	"os"
)

// Mkdir makes directories. If "createParents" is set, any parent directories
// are also created and no error is reported if a directory exists
func Mkdir(name string, mode os.FileMode, createParents bool) error {
	var err error

	if createParents {
		err = os.MkdirAll(name, mode)
	} else {
		err = os.Mkdir(name, mode)
	}

	return err
}
