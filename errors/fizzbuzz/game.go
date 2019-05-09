// © 2019 Imhotep Software LLC. All rights reserved.

package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// GameErr represents a fizzbuzz game error.
type GameErr string

func (e GameErr) Error() string {
	return string(e)
}

const (
	// Fizz number divisible by 3
	Fizz = "Fizz"
	// Buzz number divisible by 5
	Buzz = "Buzz"
	// FizzBuzz divisible by 3 and 5
	FizzBuzz = Fizz + Buzz
)

const (
	errUnderRange = GameErr("number is under range (<=0)")
	errOverRange  = GameErr("number is over range (> 20)")
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
		s = FizzBuzz
	case n%3 == 0:
		s = Fizz
	case n%5 == 0:
		s = Buzz
	default:
		s = strconv.Itoa(n)
	}

	return strconv.Itoa(n), nil
}
