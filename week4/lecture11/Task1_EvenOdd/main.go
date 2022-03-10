package main

import (
	"fmt"
	"sync"
)

func main() {

	inputs := []int{1, 17, 34, 56, 2, 8}

	var evenCh = processEven(inputs)
	fmt.Println("The even numbers are: ")
	for i := range evenCh {
		fmt.Print(i, " ")
	}

	var oddCh = processOdd(inputs)
	fmt.Println("\nThe odd numbers are: ")
	for i := range oddCh {
		fmt.Print(i, " ")
	}

}
func processEven(inputs []int) chan int {
	var evenCh = make(chan int, len(inputs))
	len := (len(inputs))
	wg := &sync.WaitGroup{}

	for i := 0; i < len; i++ {
		currentElement := inputs[i]
		wg.Add(1)
		go func(evenCh chan<- int, wg *sync.WaitGroup, currentElement int) {
			if currentElement%2 == 0 {
				evenCh <- currentElement
			}

			wg.Done()
		}(evenCh, wg, currentElement)

		wg.Wait()
	}
	close(evenCh)

	return evenCh
}
func processOdd(inputs []int) chan int {
	var oddCh = make(chan int, len(inputs))
	len := (len(inputs))
	wg := &sync.WaitGroup{}

	for i := 0; i < len; i++ {
		currentElement := inputs[i]
		wg.Add(1)
		go func(evenCh chan<- int, wg *sync.WaitGroup, currentElement int) {
			if currentElement%2 != 0 {
				oddCh <- currentElement
			}
			wg.Done()
		}(oddCh, wg, currentElement)

		wg.Wait()
	}
	close(oddCh)

	return oddCh

}
