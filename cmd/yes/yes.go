package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cd1/goreutils"
)

// YesDefaultString is the string used when an empty string is specified.
const YesDefaultString = "y"

func main() {
	var str string

	if len(os.Args) == 1 {
		str = YesDefaultString
	} else {
		str = strings.Join(os.Args[1:], " ")
	}

	ch := goreutils.Yes(str)

	for s := range ch {
		if _, err := fmt.Println(s); err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
			os.Exit(1)
		}
	}
}
