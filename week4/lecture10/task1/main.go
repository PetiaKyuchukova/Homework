package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	wc    sync.WaitGroup
	mu    sync.Mutex
	state int
	times int
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {
	cp.wc.Add(1)
	go func() {
		defer cp.wc.Done()
		for {
			cp.mu.Lock()
			if cp.state == 0 {
				fmt.Print("Foo ")
				cp.state = 1
				cp.times++
				time.Sleep(100 * time.Millisecond)
			}
			if cp.times == times {
				break
			}
			cp.mu.Unlock()

		}
		cp.wc.Done()
	}()

}

func (cp *ConcurrentPrinter) PrintBar(times int) {
	cp.wc.Add(1)

	go func(msg string) {
		defer cp.wc.Done()

		for {
			cp.mu.Lock()

			if cp.state == 1 {
				fmt.Print("Bar ")
				cp.state = 0
				cp.times++
				time.Sleep(100 * time.Millisecond)
			}

			if cp.times == times {
				break
			}

			cp.mu.Unlock()

		}
		cp.wc.Done()

	}("Bar")
}

func main() {

	times := 10
	cp := &ConcurrentPrinter{}
	cp.state = 0
	cp.times = 0
	cp.PrintFoo(times)
	cp.PrintBar(times)
	cp.wc.Wait()
}
