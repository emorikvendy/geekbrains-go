package fuzzbuzz

import (
	"io/ioutil"
	"testing"
)

func BenchmarkFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		firstSolution(ioutil.Discard)
	}
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secondSolution(ioutil.Discard)
	}
}

func BenchmarkThird(b *testing.B) {
	for i := 0; i < b.N; i++ {
		thirdSolution(ioutil.Discard)
	}
}
