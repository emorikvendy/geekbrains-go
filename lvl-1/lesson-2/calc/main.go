package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var first, second, operation string
	var result float64
	fmt.Println("Введите два чила и операцию")
	for true {
		fmt.Println("Введите первое число")
		fmt.Scanln(&first)
		if first == "q" {
			os.Exit(0)
		}
		firstNumber, err := strconv.ParseFloat(first, 64)
		if err != nil {
			fmt.Println("Ошибка: необходимо ввести число")
			continue
		}
		fmt.Println("Введите второе число")
		fmt.Scanln(&second)
		if first == "q" {
			os.Exit(0)
		}
		secondNumber, err := strconv.ParseFloat(second, 64)
		if err != nil {
			fmt.Println("Ошибка: необходимо ввести число")
			continue
		}
		fmt.Println("Введите операцию число")
		fmt.Scanln(&operation)
		switch operation {
		case "q":
			os.Exit(0)
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber + secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			if secondNumber == 0 {
				fmt.Println("Ошибка: деление на ноль")
				continue
			}
			result = firstNumber / secondNumber
		case "^":
			result = math.Pow(firstNumber, secondNumber)
		default:
			fmt.Println("Ошибка: неизвестная операция")
			continue
		}

		fmt.Printf("%f %s %f = %f\n", firstNumber, operation, secondNumber, result)
	}
}
