package wc

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Words count words given a text file path.
func Words(path string) (int, error) {
	var wc int

	f, err := os.Open(path)
	if err != nil {
		return wc, fmt.Errorf("could not open file %q: %v", path, err)
	}

	sc := bufio.NewScanner(bufio.NewReader(f))
	for sc.Scan() {
		wc += countWords(sc.Text())
	}
	if err := sc.Err(); err != nil {
		return wc, err
	}

	return wc, nil
}

func countWords(s string) int {
	wc, inWord := 0, false
	for i, r := range s {
		switch {
		case inWord && unicode.IsSpace(r):
			if inWord && i != len(s)-1 {
				wc++
			}
			inWord = false
		case unicode.IsSpace(r):
		case unicode.IsLetter(r):
			if wc == 0 {
				wc++
			}
			inWord = true
		}
	}

	return wc
}
