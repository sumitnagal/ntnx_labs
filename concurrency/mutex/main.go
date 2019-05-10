package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

type (
	bookCounts map[string]wordCounts
	wordCounts map[string]int
)

func main() {
	var (
		wg    sync.WaitGroup
		mx    sync.Mutex
		bc    = bookCounts{}
		books = []string{
			"100west.txt",
			"13chil.txt",
			"3gables.txt",
			"abbey.txt",
			"abyss.txt",
		}
		words = []string{
			"as",
			"the",
			"he",
			"she",
			"are",
		}
	)

	for _, book := range books {
		for _, word := range words {
			wg.Add(1)
			go countWord(bc, book, word, &wg, &mx)
		}
	}
	wg.Wait()

	printSummary(bc)
}

func printSummary(bc bookCounts) {
	for book, wc := range bc {
		fmt.Printf("%10s\n", book)
		for word, count := range wc {
			fmt.Printf("  %10s: %d\n", word, count)
		}
	}
}

func lock(bc bookCounts, mx *sync.Mutex, br bookRecord, f updateFn) {
	mx.Lock()
	defer mx.Unlock()
	f(bc, br)
}

type (
	updateFn   func(bc bookCounts, br bookRecord)
	bookRecord struct {
		book, word string
		count      int
	}
)

func countWord(bc bookCounts, book, word string, wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()

	br := bookRecord{
		book: book,
		word: word,
	}
	bb, err := fetch(booksURL + book)
	if err != nil {
		lock(bc, mx, br, func(bc bookCounts, br bookRecord) {
			zero := wordCounts{br.word: 0}
			if c, ok := bc[br.book]; ok {
				c[br.word] = 0
			} else {
				bc[br.book] = zero
			}
		})
		return
	}

	rx := regexp.MustCompile("(?i)" + word)
	br.count = grepCount(rx, string(bb))
	lock(bc, mx, br, func(bc bookCounts, br bookRecord) {
		if b, ok := bc[br.book]; ok {
			b[br.word] = br.count
		} else {
			bc[br.book] = wordCounts{br.word: br.count}
		}
	})
}

// Count1 returns the number of occurrence of a word in a line.
func grepCount(rx *regexp.Regexp, line string) int {
	m := rx.FindAllStringIndex(line, -1)
	return len(m)
}

// BooksURL book repo url.
const booksURL = "http://www.textfiles.com/stories/"

// Fetch books from a book repo.
func fetch(link string) ([]byte, error) {
	var res []byte
	resp, err := http.Get(link)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return res, err
	}
	res, err = ioutil.ReadAll(resp.Body)
	return res, err
}
