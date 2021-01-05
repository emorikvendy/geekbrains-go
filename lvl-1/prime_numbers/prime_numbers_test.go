package prime_numbers

import (
	"math"
	"testing"
)

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPrimeNumbers(int64(math.Pow(10, 5)))
	}
}
