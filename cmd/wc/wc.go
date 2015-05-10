package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"unicode/utf8"

	"github.com/cd1/goreutils"
)

// entry contains a WcResult and the corresponding input name (e.g. the file name, if the input is a file)
type entry struct {
	result goreutils.WcResult
	name   string
}

var (
	displayBytes         bool
	displayChars         bool
	displayLines         bool
	displayMaxLineLength bool
	displayWords         bool
)

func init() {
	flag.BoolVar(&displayBytes, "bytes", false, "print the byte counts")
	flag.BoolVar(&displayChars, "chars", false, "print the character counts")
	flag.BoolVar(&displayLines, "lines", false, "print the newline counts")
	flag.BoolVar(&displayMaxLineLength, "max-line-length", false, "print the length of the longest line")
	flag.BoolVar(&displayWords, "words", false, "print the word counts")
}

func printEntries(entries []*entry) {
	if len(entries) > 1 {
		totalEntry := &entry{
			name: "total",
		}

		// iteration #1: calculate total
		for _, e := range entries {
			totalEntry.result.Lines += e.result.Lines
			totalEntry.result.Words += e.result.Words
			totalEntry.result.Chars += e.result.Chars
			totalEntry.result.Bytes += e.result.Bytes
			totalEntry.result.MaxLineLength += e.result.MaxLineLength
		}

		entries = append(entries, totalEntry)
	}

	maxNumber := 0

	// iteration #2: calculate the number padding
	for _, e := range entries {
		maxNumber = int(math.Max(float64(e.result.Lines),
			math.Max(float64(e.result.Words),
				math.Max(float64(e.result.Chars),
					math.Max(float64(e.result.Bytes),
						math.Max(float64(e.result.MaxLineLength),
							float64(maxNumber)))))))
	}

	numberPadding := strconv.Itoa(utf8.RuneCountInString(strconv.Itoa(maxNumber)))

	if !displayBytes && !displayChars && !displayLines && !displayMaxLineLength && !displayWords {
		// default output
		displayLines = true
		displayWords = true
		displayBytes = true
	}

	// iteration #3: print the results
	for _, e := range entries {
		if displayLines {
			fmt.Printf("%"+numberPadding+"d ", e.result.Lines)
		}

		if displayWords {
			fmt.Printf("%"+numberPadding+"d ", e.result.Words)
		}

		if displayChars {
			fmt.Printf("%"+numberPadding+"d ", e.result.Chars)
		}

		if displayBytes {
			fmt.Printf("%"+numberPadding+"d ", e.result.Bytes)
		}

		if displayMaxLineLength {
			fmt.Printf("%"+numberPadding+"d ", e.result.MaxLineLength)
		}

		fmt.Printf("%v\n", e.name)
	}
}

func main() {
	flag.Parse()

	entries := make([]*entry, 0, flag.NArg())

	if flag.NArg() == 0 {
		result, err := goreutils.Wc(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
		} else {
			entries = append(entries, &entry{
				result: result,
			})
		}
	} else {
		for _, arg := range flag.Args() {
			var f *os.File
			var n string

			if arg == "-" {
				f = os.Stdin
				n = "-"
			} else {
				var err error

				if f, err = os.Open(arg); err != nil {
					fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
					continue
				}
				defer f.Close()

				n = f.Name()
			}

			result, err := goreutils.Wc(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", filepath.Base(os.Args[0]), err)
				continue
			}

			entries = append(entries, &entry{
				result: result,
				name:   n,
			})
		}
	}

	printEntries(entries)
}
