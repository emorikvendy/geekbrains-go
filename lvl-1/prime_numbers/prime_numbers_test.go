package prime_numbers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"testing"
)

func TestFindPrimeNumbers(t *testing.T) {
	var tests = []struct {
		in  int64
		out []int64
	}{
		{1, nil},
		{2, []int64{2}},
		{5, []int64{2, 3, 5}},
		{15, []int64{2, 3, 5, 7, 11, 13}},
		{30, []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{50, []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}},
	}
	for _, test := range tests {
		out := new(bytes.Buffer)
		result := findPrimeNumbers(test.in, out)
		if !reflect.DeepEqual(result, test.out) {
			t.Errorf("got %v, want %v", result, test.out)
		}
	}
}

func ExampleFindPrimeNumbers() {
	fmt.Printf("%v", findPrimeNumbers(10, os.Stdout))
	// Output: [2 3 5 7]
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPrimeNumbers(int64(math.Pow(10, 5)), ioutil.Discard)
	}
}
