package main

import (
	"fmt"
	"geekbrains-go/lvl-1/scan"
	"geekbrains-go/lvl-2/concurrency_counter"
	. "geekbrains-go/lvl-2/error_with_date"
	"os"
)

func main() {
	fmt.Println("What do you want to do? \n1 - panic\n2 - create file\n3 - run concurrency counter")
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
		concurrency_counter.Run()
	default:
		fmt.Println("Unknown code")

	}
}
