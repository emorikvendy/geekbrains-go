package main

import (
	"fmt"
	"geekbrains-go/lvl-1/scan"
	"geekbrains-go/lvl-2/concurrency_counter"
	. "geekbrains-go/lvl-2/error_with_date"
	"geekbrains-go/lvl-2/signals"
	"os"
)

const count = 10000

func main() {
	fmt.Println("What do you want to do? \n" +
		"1 - panic\n" +
		"2 - create file\n" +
		"3 - run concurrency counter with WaitGroup\n" +
		"4 - run concurrency counter\n" +
		"5 - wait for term signal\n")
	key := scan.Int64(os.Stdout, os.Stdin)
	switch key {
	case 1:
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %s", err)
			}
		}()
		fmt.Printf("%d / %d = %d", key, key-1, key/(key-1))
	case 2:
		fmt.Println("Enter filename")
		var filename string
		_, err := fmt.Scanln(&filename)
		if err != nil {
			fmt.Printf("Invalid input %v\n", err)
			return
		}
		file, err := CreateFile(filename)
		if err != nil {
			fmt.Printf("File creation error %v\n", err)
			return
		}
		defer file.Close()
		_, _ = fmt.Fprintln(file, "data")
	case 3:
		concurrency_counter.RunWG()
	case 4:
		concurrency_counter.Run()
	case 5:
		signals.WaitForTerm()
	default:
		fmt.Println("Unknown code")

	}
}
