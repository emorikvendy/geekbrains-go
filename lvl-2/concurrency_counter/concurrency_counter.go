package concurrency_counter

import (
	"fmt"
	"sync"
)

func Run() {
	wg := &sync.WaitGroup{}
	var counter int16
	var channel = make(chan struct{}, 1)
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(counter *int16, channel chan struct{}, wg *sync.WaitGroup) {
			defer func(channel chan struct{}, wg *sync.WaitGroup) {
				wg.Done()
				<-channel
			}(channel, wg)
			channel <- struct{}{}
			*counter++
		}(&counter, channel, wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
