package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cd1/goreutils"
)

var separateWithZero bool

func init() {
	flag.BoolVar(&separateWithZero, "zero", false, "separate output with NUL rather than newline")
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "%v: missing operand\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	for _, arg := range flag.Args() {
		dirname := goreutils.Dirname(arg)

		if separateWithZero {
			if _, err := fmt.Printf("%v\x00", dirname); err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
				os.Exit(1)
			}
		} else {
			if _, err := fmt.Println(dirname); err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
				os.Exit(1)
			}
		}
	}
}
