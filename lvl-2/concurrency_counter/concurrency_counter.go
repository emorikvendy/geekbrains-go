package concurrency_counter

import (
	"fmt"
	"sync"
)

func Run() {
	wg := &sync.WaitGroup{}
	var counter int16
	var channel = make(chan int8)
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(counter *int16, channel chan int8, wg *sync.WaitGroup) {
			defer func(channel chan int8, wg *sync.WaitGroup) {
				<-channel
				wg.Done()
			}(channel, wg)
			channel <- 1
			*counter++
		}(&counter, channel, wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
