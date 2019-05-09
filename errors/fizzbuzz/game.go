// © 2019 Imhotep Software LLC. All rights reserved.

package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

const (
	div3     = "Fizz"
	div5     = "Buzz"
	div3And5 = div3 + div5
)

var (
	errUnderRange = errors.New("number is under range (<=0)")
	errOverRange  = errors.New("number is over range (> 20)")
)

func main() {
	for i := 0; i <= 21; i++ {
		fmt.Printf("%02d ", i)
		if r, err := play(i); err != nil {
			fmt.Printf("Err %+v\n", err)
		} else {
			fmt.Printf("%v\n", r)
		}
	}
}

func play(n int) (string, error) {
	var s string

	if n <= 0 {
		return s, errors.Wrapf(errUnderRange, "FizzBuzz with %d", n)
	}

	if n > 20 {
		return s, errors.Wrapf(errOverRange, "FizzBuzz with %d", n)
	}

	switch {
	case n%3 == 0 && n%5 == 0:
		s = div3And5
	case n%3 == 0:
		s = div3
	case n%5 == 0:
		s = div5
	default:
		s = strconv.Itoa(n)
	}

	return s, nil
}
