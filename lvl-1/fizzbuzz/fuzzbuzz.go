package fuzzbuzz

import (
	"fmt"
	"io"
	"os"
)

func main() {
	thirdSolution(os.Stdout)
}

func firstSolution(w io.Writer) {
	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			fmt.Fprintln(w, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Fprintln(w, "Buzz")
		} else if i%3 == 0 {
			fmt.Fprintln(w, "Fizz")
		} else {
			fmt.Fprintln(w, i)
		}
	}
}

func secondSolution(w io.Writer) {
	for i := 1; i < 101; i++ {
		modulo := i % 15
		switch modulo {
		case 0:
			fmt.Fprintln(w, "FizzBuzz")
		case 3:
			fallthrough
		case 6:
			fallthrough
		case 9:
			fallthrough
		case 12:
			fmt.Fprintln(w, "Fizz")
		case 5:
			fallthrough
		case 10:
			fmt.Fprintln(w, "Buzz")
		default:
			fmt.Fprintln(w, i)
		}
	}
}

func thirdSolution(w io.Writer) {
	for i := 1; i < 101; i++ {
		if i%5 == 0 {
			if i%3 == 0 {
				fmt.Fprint(w, "Fizz")
			}
			fmt.Fprintln(w, "Buzz")

		} else if i%3 == 0 {
			fmt.Fprintln(w, "Fizz")
		} else {
			fmt.Fprintln(w, i)
		}
	}
}
