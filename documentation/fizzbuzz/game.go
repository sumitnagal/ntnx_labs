// © 2019 Imhotep Software LLC. All rights reserved.

// Package fizzbuzz implements a FizzBuzz calculator.
package fizzbuzz

import "strconv"

const (
	// Fizz number divisible by 3
	Fizz = "Fizz"
	// Buzz number divisible by 5
	Buzz = "Buzz"
	// FizzBuzz divisible by 3 and 5
	FizzBuzz = Fizz + Buzz
)

// Compute returns a FizzBuzz number based on given input.
// Returns the original number if no match.
//
// 	Returns:
//	`FizzBuzz if number is divisible by 3 and 5
// 	`Fizz if number is divisible by 3
// 	`Buzz if number is divisible by 5
// 	`number otherwise.
func Compute(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return FizzBuzz
	case n%3 == 0:
		return Fizz
	case n%5 == 0:
		return Buzz
	}

	return strconv.Itoa(n)
}
