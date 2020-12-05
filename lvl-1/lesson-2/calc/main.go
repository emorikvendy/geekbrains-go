package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var operation string
	var result float64
	fmt.Println("Введите два чила и операцию, \nДля выхода введите 'q'")
	for {
		fmt.Println("Введите первое число")
		firstNumber := scanNumber()
		fmt.Println("Введите второе число")
		secondNumber := scanNumber()
		fmt.Println("Введите операцию число")
		_, err := fmt.Scanln(&operation)
		if err != nil {
			fmt.Printf("Ошибка ввода %v\nпопробуйте ввести операцию еще раз", err)
			continue
		}
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

func scanNumber() float64 {
	var input string
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("Ошибка ввода %v\nпопробуйте ввести число еще раз", err)
			continue
		}
		if input == "q" {
			os.Exit(0)
		}
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка: необходимо ввести число")
			continue
		}
		return number
	}
}
