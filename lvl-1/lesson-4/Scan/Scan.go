package Scan

import (
	"fmt"
	"os"
	"strconv"
)

func Int64() int64 {
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
		number, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Ошибка: необходимо ввести целое число")
			continue
		}
		return number
	}
}

func Float64() float64 {
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
