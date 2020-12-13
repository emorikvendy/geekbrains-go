package Fibonacci

import (
	"math/rand"
	"testing"
)

func TestLoop(t *testing.T) {
	n := int64(rand.Intn(30))
	loop := Loop(n)
	recursive := Recursive(n)
	recursiveWithMap := RecursiveWithMap(n)
	if loop != recursive || loop != recursiveWithMap {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", loop, recursive)
	}
}
func TestRecursive(t *testing.T) {
	n := int64(rand.Intn(30))
	loop := Loop(n)
	recursive := Recursive(n)
	recursiveWithMap := RecursiveWithMap(n)
	if loop != recursive || recursive != recursiveWithMap {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", recursive, loop)
	}
}
func TestRecursiveWithMap(t *testing.T) {
	n := int64(rand.Intn(30))
	loop := Loop(n)
	recursive := Recursive(n)
	recursiveWithMap := RecursiveWithMap(n)
	if recursiveWithMap != recursive || loop != recursiveWithMap {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", recursiveWithMap, loop)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//n := int64(rand.Intn(50))
		Loop(50)
	}
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//n := int64(rand.Intn(50))
		Recursive(50)
	}
}

func BenchmarkRecursiveWithMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//n := int64(rand.Intn(50))
		RecursiveWithMap(50)
	}
}
