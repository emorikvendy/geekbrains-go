package fuzzbuzz

import "fmt"

func main() {
	thirdSolution()
}

func firstSolution() {
	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func secondSolution() {
	for i := 1; i < 101; i++ {
		modulo := i % 15
		switch modulo {
		case 0:
			fmt.Println("FizzBuzz")
		case 3:
			fallthrough
		case 6:
			fallthrough
		case 9:
			fallthrough
		case 12:
			fmt.Println("Fizz")
		case 5:
			fallthrough
		case 10:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func thirdSolution() {
	for i := 1; i < 101; i++ {
		if i%5 == 0 {
			if i%3 == 0 {
				fmt.Print("Fizz")
			}
			fmt.Println("Buzz")

		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
