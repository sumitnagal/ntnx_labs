// © 2019 Imhotep Software LLC. All rights reserved.

package wc

import (
	"index/suffixarray"
	"regexp"
	"strings"
)

// Count returns the number of occurrence of a word in a line.
func Count(word, line string) int {
	rx := regexp.MustCompile(strings.ToLower(word))
	m := rx.FindAllStringIndex(strings.ToLower(line), -1)
	return len(m)
}

// Count1 returns the number of occurrence of a word in a line.
func Count1(rx *regexp.Regexp, line string) int {
	m := rx.FindAllStringIndex(line, -1)
	return len(m)
}

// Count2 returns the number of occurrence of a word in a line.
func Count2(rx *regexp.Regexp, line string) int {
	index := suffixarray.New([]byte(line))
	m := index.FindAllIndex(rx, -1)
	return len(m)
}
