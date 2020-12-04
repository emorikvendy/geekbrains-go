package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBubble(t *testing.T) {
	slice := randomInt64Slice()
	sorted := Bubble(slice)
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	if !Equal(sorted, slice) {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", sorted, slice)
	}
}
func TestInsertion(t *testing.T) {
	slice := randomInt64Slice()
	sorted := Insertion(slice)
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	if !Equal(sorted, slice) {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", sorted, slice)
	}
}
func TestBuiltIn(t *testing.T) {
	slice := randomInt64Slice()
	sorted := BuiltIn(slice)
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	if !Equal(sorted, slice) {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", sorted, slice)
	}
}

func randomInt64Slice() []int64 {
	rand.Seed(time.Now().UnixNano())
	size := 1000
	slice := make([]int64, size, size)
	for i := 0; i < len(slice); i++ {
		slice[i] = int64(rand.Intn(100 * size))
	}
	return slice
}

func Equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := randomInt64Slice()
		Bubble(slice)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := randomInt64Slice()
		Insertion(slice)
	}
}

func BenchmarkBuiltIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := randomInt64Slice()
		BuiltIn(slice)
	}
}
