package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cd1/goreutils"
)

var (
	multipleArgs     bool
	separateWithZero bool
	suffix           string
)

func init() {
	flag.BoolVar(&multipleArgs, "multiple", false, "support multiple arguments and treat each as a NAME")
	flag.StringVar(&suffix, "suffix", "", "remove a trailing SUFFIX")
	flag.BoolVar(&separateWithZero, "zero", false, "separate output with NUL rather than newline")
}

func main() {
	flag.Parse()

	// --suffix implies --multiple
	if suffix != "" {
		multipleArgs = true
	}

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "%v: missing operand\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	bases := make([]string, 0, flag.NArg())

	if multipleArgs {
		for _, arg := range flag.Args() {
			bases = append(bases, goreutils.Basename(arg, suffix))
		}
	} else {
		if flag.NArg() > 2 {
			fmt.Fprintf(os.Stderr, "%v: extra operand '%v'\n", filepath.Base(os.Args[0]), flag.Arg(2))
			os.Exit(1)
		}

		if flag.NArg() == 2 {
			suffix = flag.Arg(1)
		}

		bases = append(bases, goreutils.Basename(flag.Arg(0), suffix))
	}

	if separateWithZero {
		if _, err := fmt.Print(strings.Join(bases, "\x00")); err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
			os.Exit(1)
		}
	} else {
		if _, err := fmt.Println(strings.Join(bases, "\n")); err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
			os.Exit(1)
		}
	}
}
