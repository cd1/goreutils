package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cd1/goreutils"
)

func main() {
	if _, err := fmt.Println(goreutils.Arch()); err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}
