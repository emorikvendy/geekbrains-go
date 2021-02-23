package calc

import (
	"fmt"
	"geekbrains-go/lvl-1/scan"
	"math"
	"os"
)

func calc() {
	var operation string
	var result float64
	fmt.Println("Введите два чила и операцию, \nДля выхода введите 'q'")
	for {
		fmt.Println("Введите первое число")
		firstNumber := scan.Float64(os.Stdout, os.Stdin)
		fmt.Println("Введите второе число")
		secondNumber := scan.Float64(os.Stdout, os.Stdin)
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
