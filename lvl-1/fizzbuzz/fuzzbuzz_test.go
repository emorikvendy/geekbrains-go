package fuzzbuzz

import (
	"testing"
)

func BenchmarkFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstSolution()
	}
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secondSolution()
	}
}

func BenchmarkThird(b *testing.B) {
	for i := 0; i < b.N; i++ {
		thirdSolution()
	}
}
