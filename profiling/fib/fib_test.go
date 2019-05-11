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

func TestFibRec(t *testing.T) {
	for _, u := range uu {
		assert.Equal(t, u.e, fib.ComputeRec(u.n))
	}
}

func TestComputeLinear(t *testing.T) {
	for _, u := range uu {
		assert.Equal(t, u.e, fib.ComputeLinear(u.n))
	}
}

func BenchmarkFibRec(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fib.ComputeRec(20)
	}
}

func BenchmarkFibCompute(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fib.ComputeLinear(20)
	}
}
