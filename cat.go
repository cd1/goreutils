package goreutils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Cat concatenates files and prints on an io.Reader.
func Cat(inputs []io.Reader, out io.Writer, showTabs, number, numberNonBlank, showEnds bool) error {
	lineNumber := 1

	for _, in := range inputs {
		scanner := bufio.NewScanner(in)

		for scanner.Scan() {
			text := scanner.Text()

			if showTabs {
				text = strings.Replace(text, "\t", "^I", -1)
			}

			if numberNonBlank {
				if text != "" {
					if _, err := fmt.Fprintf(out, "%6d  ", lineNumber); err != nil {
						return err
					}
					lineNumber++
				}
			} else if number {
				if _, err := fmt.Fprintf(out, "%6d  ", lineNumber); err != nil {
					return err
				}
				lineNumber++
			}

			if _, err := fmt.Fprint(out, text); err != nil {
				return err
			}

			if showEnds {
				if _, err := fmt.Fprint(out, "$"); err != nil {
					return err
				}
			}

			if _, err := fmt.Fprintln(out); err != nil {
				return err
			}
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}
