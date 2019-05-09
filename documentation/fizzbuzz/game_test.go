// © 2019 Imhotep Software LLC. All rights reserved.

package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/gopherland/labs/documentation/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

var Output string

func TestCompute(t *testing.T) {
	uu := map[string]struct{
		n int
		e string
	} {
		"zero":     {0, fizzbuzz.FizzBuzz},
		"one":      {1, "1"},
		"fizz":     {3, fizzbuzz.Fizz},
		"four":     {4,  "4"},
		"buzz":     {5,  fizzbuzz.Buzz},
		"fizzbuzz": {15, fizzbuzz.FizzBuzz},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, fizzbuzz.Compute(u.n))
		})
	}
}

func BenchmarkCompute(b *testing.B) {
	var out string
	for i := 0; i < b.N; i++ {
		out = fizzbuzz.Compute(i)
	}
	Output = out
}

// Returns `FizzBuzz if number is div by 3 and 5
func ExampleCompute_DivisibleBy3And5() {
	fmt.Println(fizzbuzz.Compute(15))
	// Output:
	// FizzBuzz
}

// Returns `Fizz if number is div by 3
func ExampleCompute_DivisibleBy3() {
	fmt.Println(fizzbuzz.Compute(3))
	// Output:
	// Fizz
}

// Returns `Buzz if number is div by 5
func ExampleCompute_DivisibleBy5() {
	fmt.Println(fizzbuzz.Compute(5))
	// Output:
	// Buzz
}
