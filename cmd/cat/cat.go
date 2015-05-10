package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cd1/goreutils"
)

var (
	number         bool
	numberNonBlank bool
	showEnds       bool
	showTabs       bool
	u              bool
)

func init() {
	flag.BoolVar(&numberNonBlank, "b", false, "number nonempty output lines, overrides -n")
	flag.BoolVar(&number, "n", false, "number all output lines")
	flag.BoolVar(&showEnds, "E", false, "display $ at end of each line")
	flag.BoolVar(&showTabs, "T", false, "display TAB characters as ^I")
	flag.BoolVar(&u, "u", false, "(ignored)")
}

func main() {
	flag.Parse()

	exitCode := 0

	if flag.NArg() == 0 {
		if err := goreutils.Cat([]io.Reader{os.Stdin}, os.Stdout, showTabs, number, numberNonBlank, showEnds); err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
			os.Exit(1)
		}
	} else {
		readers := make([]io.Reader, 0, flag.NArg())

		for _, arg := range flag.Args() {

			if arg == "-" {
				readers = append(readers, os.Stdin)
			} else {
				file, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
					exitCode = 1
					continue
				}
				defer file.Close()

				readers = append(readers, file)
			}
		}

		if err := goreutils.Cat(readers, os.Stdout, showTabs, number, numberNonBlank, showEnds); err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
			exitCode = 1
		}

		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
}
