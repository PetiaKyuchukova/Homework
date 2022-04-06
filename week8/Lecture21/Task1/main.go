package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type Article struct {
	Id     int    `json:"-"`
	Tittle string `json:"title"`
	Score  int    `json:"score"`
}
type News struct {
	Articles []Article `json:"top_stories"`
}

func HandlerNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup
		var m sync.Mutex
		var news News

		news.Articles = make([]Article, 10, 10)

		ids := GetIDs()

		for i := 0; i < len(ids); i++ {
			news.Articles[i].Id = ids[i]
		}

		for i, v := range news.Articles {
			wg.Add(1)
			go func(i int, v Article) {
				m.Lock()
				GetDetails(&news.Articles[i])
				m.Unlock()
				wg.Done()
			}(i, v)

		}
		wg.Wait()

		empData, err := json.MarshalIndent(news, "", "\t")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(empData))
		w.Write(empData)
	}

}
func GetDetails(a *Article) {

	path := ("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(a.Id) + ".json?print=pretty")
	urlA, err := url.Parse(path)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Get(urlA.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(body, &a)
	if err1 != nil {
		fmt.Println(err1)
	}

}
func GetIDs() [10]int {
	var IDs [10]int

	urlA, err := url.Parse("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(urlA.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(body, &IDs)
	if err1 != nil {
		fmt.Println(err1)
	}
	return IDs

}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/top", HandlerNews())
	http.ListenAndServe(":9000", router)

}
