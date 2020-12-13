package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var input string
	var err error
	fmt.Println("Ввведите целое число")
	_, err = fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Ошибка ввода %v\n", err)
		os.Exit(0)
	}
	maxNumber, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Printf("Некорректный ввод %v\n", err)
		os.Exit(0)
	}
	primeNumbers := findPrimeNumbers(maxNumber)
	fmt.Println(primeNumbers)
}

func findPrimeNumbers(maxNumber int64) []int64 {

	if maxNumber <= 1 {
		fmt.Println("Минимальное простое число равно 2")
		return nil
	}
	sliceSize := math.Round(float64(maxNumber)/math.Log(float64(maxNumber))) + 1
	var primeNumbers = make([]int64, sliceSize, sliceSize)
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
