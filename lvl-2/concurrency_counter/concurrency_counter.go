package concurrency_counter

import (
	"fmt"
	"sync"
	"time"
)

const count = 1000

func RunWG() {
	wg := &sync.WaitGroup{}
	var counter int16
	var channel = make(chan struct{}, 1)
	defer close(channel)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-channel
			}()
			channel <- struct{}{}
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func Run() {
	var (
		counter       int16
		lockChannel   = make(chan struct{}, 1)
		finishChannel = make(chan struct{}, count)
	)
	for i := 0; i < count; i++ {
		finishChannel <- struct{}{}
		go func() {
			defer func() {
				<-lockChannel
				<-finishChannel
			}()
			lockChannel <- struct{}{}
			counter++
		}()
	}

	for true {
		if len(finishChannel) == 0 {
			break
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Println(counter)
}
