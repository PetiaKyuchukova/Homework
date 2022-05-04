package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type ArticleService struct {
	Url string
}
type Article struct {
	Id    int    `json:"-"`
	Title string `json:"title"`
	Score int    `json:"score"`
	Time  string `json:"-"`
}

func NewArticleService(path string) *ArticleService {
	return &ArticleService{Url: path}
}
func (as *ArticleService) GetDetails(a *Article) {

	serverUrl := (as.Url + "/item/" + strconv.Itoa(a.Id) + ".json?print=pretty")
	urlA, err := url.Parse(serverUrl)
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
func (as *ArticleService) GetIDs() [10]int {
	var IDs [10]int
	var body []byte

	urlA, err := url.Parse(as.Url + "/topstories.json?print=pretty")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(urlA.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	err1 := json.Unmarshal(body, &IDs)
	if err1 != nil {
		fmt.Println(err1)
	}
	return IDs

}
