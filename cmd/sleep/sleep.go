package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unicode"

	"github.com/cd1/goreutils"
)

const defaultUnit = "s"

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "%v: missing operand\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	duration := time.Duration(0)

	for _, arg := range os.Args[1:] {
		if unicode.IsDigit(rune(arg[len(arg)-1])) {
			arg += defaultUnit
		}

		dur, err := time.ParseDuration(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
			os.Exit(1)
		}

		duration += dur
	}

	goreutils.Sleep(duration)
}
