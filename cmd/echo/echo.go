package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cd1/goreutils"
)

var suppressNewLine bool

func init() {
	flag.BoolVar(&suppressNewLine, "n", false, "do not output the trailing newline")
}

func main() {
	flag.Parse()

	var err error

	str := goreutils.Echo(strings.Join(flag.Args(), " "))

	if suppressNewLine {
		_, err = fmt.Print(str)
	} else {
		_, err = fmt.Println(str)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}
