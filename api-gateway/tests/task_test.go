package tests

import (
	"testing"
)

func TestFib3(T *testing.T) {
	Fib3(10)
}

func BenchmarkFib3(b *testing.B) {
	for n:=0;n<b.N;n++ {
		Fib3(20)
	}
}
