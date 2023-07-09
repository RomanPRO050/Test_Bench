package util

import "testing"

func BenchmarkFib(b *testing.B) {
	Fib(35)
}

func BenchmarkMakeSlice(b *testing.B) {
	MakeSlice(3000000)
}

func BenchmarkFibStraight(b *testing.B) {
	FibStraight(35)
}
