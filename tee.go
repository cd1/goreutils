package goreutils

import (
	"io"
)

// Tee reads from one input and writes to multiple outputs.
func Tee(in io.Reader, out []io.Writer) error {
	_, err := io.Copy(io.MultiWriter(out...), in)

	return err
}
