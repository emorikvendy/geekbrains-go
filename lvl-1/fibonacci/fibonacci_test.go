package fibonacci

import (
	"math/rand"
	"testing"
)

var testValues = []int64{-25, -2, -1, 0, 1, 10, 30}

func TestLoop(t *testing.T) {
	for _, n := range testValues {
		loop := Loop(n)
		recursive := Recursive(n)
		recursiveWithMap := RecursiveWithMap(n)
		if loop != recursive || loop != recursiveWithMap {
			t.Errorf("results not match when n:=%d\nGot:\n%v\nExpected:\n%v", n, loop, recursive)
		}
	}
}
func TestRecursive(t *testing.T) {
	for _, n := range testValues {
		loop := Loop(n)
		recursive := Recursive(n)
		recursiveWithMap := RecursiveWithMap(n)
		if loop != recursive || recursive != recursiveWithMap {
			t.Errorf("results not match when n:=%d\nGot:\n%v\nExpected:\n%v", n, recursive, loop)
		}
	}
}
func TestRecursiveWithMap(t *testing.T) {
	for _, n := range testValues {
		loop := Loop(n)
		recursive := Recursive(n)
		recursiveWithMap := RecursiveWithMap(n)
		if recursiveWithMap != recursive || loop != recursiveWithMap {
			t.Errorf("results not match when n:=%d\nGot:\n%v\nExpected:\n%v", n, recursiveWithMap, loop)
		}
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := int64(rand.Intn(61)) - 30
		Loop(n)
	}
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := int64(rand.Intn(61)) - 30
		Recursive(n)
	}
}

func BenchmarkRecursiveWithMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := int64(rand.Intn(61)) - 30
		RecursiveWithMap(n)
	}
}
