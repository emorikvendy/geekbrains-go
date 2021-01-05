package main

import (
	"fmt"
	"geekbrains-go/lvl-1/fibonacci"
	"geekbrains-go/lvl-1/scan"
	"os"
)

func main() {
	fmt.Println("Введите номер числа Фибоначчи, которое хотите получить")
	n := scan.Int64(os.Stdout)
	fmt.Println("Loop result: ", fibonacci.Loop(n))
	//fmt.Println("Recursive result: ", Fibonacci.Recursive(n))
	fmt.Println("RecursiveWithMap result: ", fibonacci.RecursiveWithMap(n))
}
