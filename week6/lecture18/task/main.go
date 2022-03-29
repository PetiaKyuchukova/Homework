package primes

import (
	"sync"
	"time"
)

func main() {

}
func goPrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for k := 2; k < n; k++ {
		wg.Add(1)
		go func(k int) {
			for i := 2; i < n; i++ {
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i {

						mu.Lock()
						res = append(res, k)
						mu.Unlock()

					}
					wg.Done()
					break
				}
			}
		}(k)

	}
	wg.Wait()
	return res
}
func primesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}
