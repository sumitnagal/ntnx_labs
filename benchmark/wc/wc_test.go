package wc_test

import (
	"regexp"
	"testing"

	"github.com/gopherland/level2/labs/wc"
	"gotest.tools/assert"
)

func TestWc(t *testing.T) {
	uu := map[string]struct {
		word, line string
		count      int
	}{
		"matches": {
			"fox",
			"The brown red fox, is a really cool FOX!",
			2,
		},
		"upper": {
			"Fox",
			"The brown red fox, is a really cool FOXY!",
			2,
		},
		"zero": {
			"Bob",
			"The brown red fox, is a really cool FOX!",
			0,
		},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.count, wc.Count(u.word, u.line))
		})
	}
}

func TestWc1(t *testing.T) {
	uu := map[string]struct {
		word, line string
		count      int
	}{
		"matches": {
			"fox",
			"The brown red fox, is a really cool FOX!",
			2,
		},
		"upper": {
			"Fox",
			"The brown red fox, is a really cool FOXY!",
			2,
		},
		"zero": {
			"Bob",
			"The brown red fox, is a really cool FOX!",
			0,
		},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			rx := regexp.MustCompile("(?i)" + u.word)
			assert.Equal(t, u.count, wc.Count1(rx, u.line))
		})
	}
}

func TestWc2(t *testing.T) {
	uu := map[string]struct {
		word, line string
		count      int
	}{
		"matches": {
			"fox",
			"The brown red fox, is a really cool FOX!",
			2,
		},
		"upper": {
			"Fox",
			"The brown red fox, is a really cool FOXY!",
			2,
		},
		"zero": {
			"Bob",
			"The brown red fox, is a really cool FOX!",
			0,
		},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			rx := regexp.MustCompile("(?i)" + u.word)
			assert.Equal(t, u.count, wc.Count2(rx, u.line))
		})
	}
}

func BenchmarkCount1(b *testing.B) {
	w, line := "fox", "The brown red fox, is a really cool FOX!"
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		wc.Count(w, line)
	}
}

func BenchmarkCount2(b *testing.B) {
	w, line := "fox", "The brown red fox, is a really cool FOX!"
	rx := regexp.MustCompile("(?i)" + w)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		wc.Count1(rx, line)
	}
}

func BenchmarkCount3(b *testing.B) {
	w, line := "fox", "The brown red fox, is a really cool FOX!"
	rx := regexp.MustCompile("(?i)" + w)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		wc.Count2(rx, line)
	}
}
