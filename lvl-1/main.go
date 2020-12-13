package main

import (
	"fmt"
	"geekbrains-go/lvl-1/lesson-4/Fibonacci"
	"geekbrains-go/lvl-1/lesson-4/Scan"
)

func main() {
	fmt.Println("Введите номер числа Фибоначчи, которое хотите получить")
	n := Scan.Int64()
	fmt.Println("Loop result: ", Fibonacci.Loop(n))
	//fmt.Println("Recursive result: ", Fibonacci.Recursive(n))
	fmt.Println("RecursiveWithMap result: ", Fibonacci.RecursiveWithMap(n))
}
