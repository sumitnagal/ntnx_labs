package fib_test

import (
	"testing"

	"github.com/gopherland/ntnx_labs/performance/fib"
	"github.com/stretchr/testify/assert"
)

var uu = []struct {
	n, e int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
}

func TestFibCompute(t *testing.T) {
	for _, u := range uu {
		assert.Equal(t, u.e, fib.Compute(u.n))
	}
}

func BenchmarkFibCompute(b *testing.B) {
	// !!YOUR_CODE!!
}
