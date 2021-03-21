package main

import (
	"fmt"
	"geekbrains-go/lvl-1/scan"
	"geekbrains-go/lvl-2/concurrency_counter"
	. "geekbrains-go/lvl-2/error_with_date"
	"geekbrains-go/lvl-2/go_parser"
	"geekbrains-go/lvl-2/go_sched"
	"geekbrains-go/lvl-2/signals"
	"os"
	"runtime"
	"runtime/trace"
)

const count = 10000

func main() {
	fmt.Printf("What do you want to do? \n" +
		"1 - panic\n" +
		"2 - create file\n" +
		"3 - run concurrency counter with WaitGroup\n" +
		"4 - run concurrency counter\n" +
		"5 - run concurrency counter with Mutex\n" +
		"6 - wait for term signal\n" +
		"7 - trace concurrency counter with Mutex\n" +
		"8 - run concurrency counter with race\n" +
		"9 - run long process with GoSched\n" +
		"10 - run long process without GoSched\n" +
		"11 - run go parser\n")
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
		concurrency_counter.RunMutex()
	case 6:
		signals.WaitForTerm()
	case 7:
		trace.Start(os.Stderr)
		defer trace.Stop()
		concurrency_counter.RunMutex()
	case 8:
		concurrency_counter.RunRace()
	case 9:
		runtime.GOMAXPROCS(1)
		trace.Start(os.Stderr)
		defer trace.Stop()
		go_sched.Run()
	case 10:
		runtime.GOMAXPROCS(1)
		trace.Start(os.Stderr)
		defer trace.Stop()
		go_sched.RunWithoutGoShed()
	case 11:
		sep := string(os.PathSeparator)
		go_parser.Run("lvl-2"+sep+"concurrency_counter"+sep+"concurrency_counter.go", "RunWG")
	default:
		fmt.Println("Unknown code")

	}
}
