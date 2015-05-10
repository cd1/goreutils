package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cd1/goreutils"
)

var (
	mode    int
	parents bool
	verbose bool
)

func init() {
	flag.IntVar(&mode, "mode", 0777, "set file mode")
	flag.BoolVar(&parents, "parents", false, "no error if existing, make parent directories as needed")
	flag.BoolVar(&verbose, "verbose", false, "print a message for each created directory")
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "%v: missing operand\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	exitCode := 0

	for _, arg := range flag.Args() {
		if err := goreutils.Mkdir(arg, os.FileMode(mode), parents); err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
			exitCode = 1
			continue
		}

		if verbose {
			if _, err := fmt.Printf("%v: created directory '%v'\n", filepath.Base(os.Args[0]), arg); err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
				os.Exit(1)
			}
		}
	}

	os.Exit(exitCode)
}
