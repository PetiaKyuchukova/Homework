package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type BufferedContext struct {
	cnt         context.Context
	bufferSize  int
	currentSize int
	buffer      chan string
	wg          sync.WaitGroup
	mutex       sync.Mutex
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	currentSize := 0
	buffer := make(chan string, bufferSize)
	buffContext := BufferedContext{ctx, bufferSize, currentSize, buffer, sync.WaitGroup{}, sync.Mutex{}}

	go func() {
		time.Sleep(timeout) 
		fmt.Println("Timeout!")
		cancel()
	}()

	return &buffContext
}
func (bc *BufferedContext) DoneB() <-chan struct{} {
	fmt.Println("DoneB")
	return bc.cnt.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	fmt.Println("Run")

	go func() {
		for {
			if len(bc.buffer) == cap(bc.buffer) {
				fmt.Println("Buffer full!")
				bc.DoneB()
				return
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()

	bc.wg.Add(1)
	go func() {
		fn(bc.cnt, bc.buffer)
		bc.wg.Done()
	}()
	bc.wg.Wait()
}

func main() {
	bufferedCtx := NewBufferedContext(time.Second, 10)

	bufferedCtx.Run(
		func(bufferedCtx context.Context, buffer chan string) {
			for {
				select {
				case <-bufferedCtx.Done():
					return
				case buffer <- "bar":
					time.Sleep(time.Millisecond * 200) // try different values here
					fmt.Println("bar")
				}
			}
		})
}
