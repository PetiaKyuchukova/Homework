package main

import (
	//"crypto/rand"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

func fetchURLs(urls []string, concureency int) {

	processQueue := make(chan string, concureency)
	var wg sync.WaitGroup

	for _, urlToProcess := range urls {
		wg.Add(1)

		processQueue <- urlToProcess

		go func(url string) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(2000)) * time.Microsecond)
			pingURL(<-processQueue)

		}(urlToProcess)

	}
	wg.Wait()
	close(processQueue)

}

func main() {

	countPtr := flag.Int("c", 2, "crawler count")
	flag.Parse()

	count := *countPtr
	filePaths := flag.Args()

	if len(filePaths) > count {
		log.Print("The crawler count is smaller than given argoments!")
		flag.PrintDefaults()
		flag.Usage()
		return
	} else {
		fetchURLs(filePaths, count)
	}

}
