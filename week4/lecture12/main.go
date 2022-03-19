package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	channel := make(chan string)
	channelBuff := make(chan string, bufferLimit)

	go func() {
		for {
			channel <- data
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			time.Sleep(clearInterval)
			for i := 0; i < bufferLimit; i++ {
				channelElement := <-channel
				channelBuff <- channelElement
			}
		}
	}()

	return channelBuff
}

func main() {
	out := generateThrottled("foo", 2, time.Second)

	for f := range out {
		log.Println(f)
	}
}
