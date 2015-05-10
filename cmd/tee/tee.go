package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/cd1/goreutils"
)

var (
	appendToFile     bool
	ignoreInterrupts bool
)

func init() {
	flag.BoolVar(&appendToFile, "append", false, "append to the given files, do not overwrite")
	flag.BoolVar(&ignoreInterrupts, "ignore-interrupts", false, "ignore interrupt signals")
}

func main() {
	flag.Parse()

	if ignoreInterrupts {
		sigintCh := make(chan os.Signal)
		go func() {
			for {
				<-sigintCh
			}
		}()

		signal.Notify(sigintCh, os.Interrupt)
	}

	outputs := []io.Writer{os.Stdout}
	exitCode := 0

	for _, arg := range flag.Args() {
		if arg == "-" {
			outputs = append(outputs, os.Stdout)
		} else {
			flags := os.O_CREATE | os.O_WRONLY

			if appendToFile {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}

			f, err := os.OpenFile(arg, flags, 0666)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
				exitCode = 1
				continue
			}
			defer f.Close()

			outputs = append(outputs, f)
		}
	}

	if err := goreutils.Tee(os.Stdin, outputs); err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
		exitCode = 1
	}

	os.Exit(exitCode)
}
