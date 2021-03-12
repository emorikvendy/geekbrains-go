package signals

import (
	"fmt"
	"geekbrains-go/lvl-2/concurrency_counter"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitForTerm() {
	var (
		sigs        = make(chan os.Signal, 1)
		done        = make(chan bool, 1)
		itterations int64
	)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	for true {
		select {
		case <-done:
			time.Sleep(900 * time.Millisecond)
			fmt.Printf("%d itterations done\n", itterations)
			os.Exit(0)
		default:
			concurrency_counter.Run()
			itterations++
		}
	}
}
