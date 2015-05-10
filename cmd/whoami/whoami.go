package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cd1/goreutils"
)

func main() {
	me, err := goreutils.Whoami()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}

	if _, err := fmt.Println(me); err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}
