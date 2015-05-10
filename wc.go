package goreutils

import (
	"bufio"
	"io"
	"math"
	"strings"
	"unicode/utf8"
)

// WcResult contains the result data of the function "Wc".
type WcResult struct {
	Lines         int
	Words         int
	Chars         int
	Bytes         int
	MaxLineLength int
}

// Wc counts newline, word, char and byte counts for an input, as well as the
// maximum line length.
func Wc(in io.Reader) (WcResult, error) {
	linesCount := 0
	wordsCount := 0
	charsCount := 0
	bytesCount := 0
	maxLineLength := 0

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		text := scanner.Text()

		w := len(strings.Fields(text))
		c := utf8.RuneCountInString(text)
		b := len(text)

		linesCount++
		wordsCount += w
		charsCount += c
		bytesCount += b
		maxLineLength = int(math.Max(float64(maxLineLength), float64(c)))
	}
	if err := scanner.Err(); err != nil {
		return WcResult{}, err
	}

	r := WcResult{
		Lines:         linesCount,
		Words:         wordsCount,
		Chars:         charsCount,
		Bytes:         bytesCount,
		MaxLineLength: maxLineLength,
	}

	return r, nil
}
