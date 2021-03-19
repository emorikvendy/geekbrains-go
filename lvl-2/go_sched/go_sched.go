package go_sched

import (
	"runtime"
	"sync"
)

const count = 10

func Run() {
	wg := &sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			for j := 1; j < 1e6; j++ {
				if j%1e5 == 0 {
					runtime.Gosched()
				}
			}
		}()
	}
	wg.Wait()
}

func RunWithoutGoShed() {
	wg := &sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			for j := 1; j < 1e6; j++ {
			}
		}()
	}
	wg.Wait()
}
