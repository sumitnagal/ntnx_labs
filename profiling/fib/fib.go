package fib

// ComputeRec computes Fibonacci number recursively.
func ComputeRec(n int) int {
	if n < 2 {
		return n
	}
	return ComputeRec(n-2) + ComputeRec(n-1)
}

// ComputeLinear computes Fibonacci number in line.
func ComputeLinear(n int) int {
	if n < 2 {
		return n
	}
	p1, p2 := 0, 1
	for i := 0; i < n-1; i++ {
		p1, p2 = p2, p1+p2
	}
	return p2
}
