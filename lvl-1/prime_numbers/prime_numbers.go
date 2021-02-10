package prime_numbers

import (
	"fmt"
	"geekbrains-go/lvl-1/scan"
	"io"
	"math"
	"os"
)

func run() {
	fmt.Println("Ввведите целое число")
	maxNumber := scan.Int64(os.Stdout, os.Stdin)
	primeNumbers := findPrimeNumbers(maxNumber, os.Stdout)
	fmt.Println(primeNumbers)
}

func findPrimeNumbers(maxNumber int64, w io.Writer) []int64 {

	if maxNumber <= 1 {
		_, err := fmt.Fprintln(w, "Минимальное простое число равно 2")
		if err != nil {
			panic("Ошибка вывода")
		}
		return nil
	}
	sliceSize := int64(math.Round(float64(maxNumber)/math.Log(float64(maxNumber)))) + 1
	var primeNumbers = make([]int64, 0, sliceSize)
	primeNumbers = append(primeNumbers, 2)
	for i := int64(3); i <= maxNumber; i += 2 {
		isComposite := false
		sqrt := int64(math.Sqrt(float64(i)))
		for _, prime := range primeNumbers {
			if prime > sqrt {
				break
			}
			if i%prime == 0 {
				isComposite = true
				break
			}
		}
		if !isComposite {
			primeNumbers = append(primeNumbers, i)
		}
	}

	return primeNumbers
}
