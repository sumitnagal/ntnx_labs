package greeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Output string

func TestGreeter1(t *testing.T) {
	assert.Equal(t, "Hello, Fernand! You are 42 old today...", greeter1("Fernand", 42))
}

func TestGreeter2(t *testing.T) {
	assert.Equal(t, "Hello, Fernand! You are 42 old today...", greeter2("Fernand", 42))
}

func TestGreeter3(t *testing.T) {
	assert.Equal(t, "Hello, Fernand! You are 42 old today...", greeter3("Fernand", 42))
}

func TestGreeter4(t *testing.T) {
	assert.Equal(t, "Hello, Fernand! You are 42 old today...", greeter4("Fernand", 42))
}

func BenchmarkGreeter1(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter1(n, a)
	}
}

func BenchmarkGreeter2(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter2(n, a)
	}
}

func BenchmarkGreeter3(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter3(n, a)
	}
}

func BenchmarkGreeter4(b *testing.B) {
	n, a := "Fernand", 30
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greeter4(n, a)
	}
}
