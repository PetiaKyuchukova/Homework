package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

type Article struct {
	Id    int    `json:"-"`
	Title string `json:"title"`
	Score int    `json:"score"`
	Time  string `json:"-"`
}
type News struct {
	PageTitle string    `json:"PageTitle"`
	Articles  []Article `json:"top_stories"`
}

func HNmock() []byte {
	return []byte(`{
		"PageTitle": "Top 10 stories",
		"top_stories": [
			{
				"title": "Zas Editor",
				"score": 390
			},
			{
				"title": "Buffereditor  Code and Text Editor for iPhone, iPad and iPad Pro",
				"score": 17
			},
			{
				"title": "You may not need Cloudflare Tunnel. Linux is fine",
				"score": 66
			},
			{
				"title": "PiGlass v2: A wearable Pi Zero 2",
				"score": 64
			},
			{
				"title": "Is the DTS vs. Dolby war effectively over?",
				"score": 28
			},
			{
				"title": "Ask HN: Hybrid/Remote software team rituals",
				"score": 38
			},
			{
				"title": "1.5-Inch Hand-Carved Agate from 1500 B.C. Shows Sub-Millimeter Details",
				"score": 176
			},
			{
				"title": "Reading academic computer science papers",
				"score": 130
			},
			{
				"title": "Firefox DNS-over-HTTPS",
				"score": 113
			},
			{
				"title": "David Deutschs Constructor Theory",
				"score": 64
			}
		]
	}`)
}
func HandlerNewsHTMLTemplate(path string) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	return func(w http.ResponseWriter, r *http.Request) {

		var wg sync.WaitGroup
		var m sync.Mutex
		var news News
		news.PageTitle = "Top 10 stories"
		news.Articles = CheckDB()
		if len(news.Articles) == 10 {
			tmpl.Execute(w, news)
			log.Print("FromDB")
		} else {
			log.Print("else")

			news.Articles = make([]Article, 10, 10)

			ids := GetIDs(path)

			for i := 0; i < len(ids); i++ {
				news.Articles[i].Id = ids[i]
				news.Articles[i].Time = time.Now().Format("2006-01-02 15:04:05")
			}

			for i, v := range news.Articles {
				wg.Add(1)
				go func(i int, v Article) {
					m.Lock()
					GetDetails(&news.Articles[i], path)
					m.Unlock()
					wg.Done()
				}(i, v)
			}

			wg.Wait()

			tmpl.Execute(w, news)

			db, err := sql.Open("sqlite", "data.db")
			if err != nil {
				log.Fatal(err)
			}

			defer db.Close()

			Delete(db)

			for _, v := range news.Articles {
				Insert(db, v.Id, v.Title, v.Score, v.Time)
			}
		}

	}
}
func HandlerNewsMarshal(path string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var wg sync.WaitGroup
		var m sync.Mutex
		var news News
		news.PageTitle = "Top 10 stories"

		news.Articles = CheckDB()

		if len(news.Articles) == 10 {
			empData, err := json.MarshalIndent(news, "", "\t")
			if err != nil {
				fmt.Println(err)
			}

			w.Write(empData)

		} else {
			if path == "https://hacker-news.firebaseio.com/v0" {
				ids := GetIDs(path)

				news.Articles = make([]Article, 10, 10)

				for i := 0; i < len(ids); i++ {

					news.Articles[i].Id = ids[i]
					news.Articles[i].Time = time.Now().Format("2006-01-02 15:04:05")

				}

				for i, v := range news.Articles {
					wg.Add(1)
					go func(i int, v Article) {
						m.Lock()
						GetDetails(&news.Articles[i], path)
						m.Unlock()
						wg.Done()
					}(i, v)

				}
				wg.Wait()

				empData, err := json.MarshalIndent(news, "", "\t")
				if err != nil {
					log.Fatal(err)
				}

				w.Write(empData)

				db, err := sql.Open("sqlite", "data.db")
				if err != nil {
					log.Fatal(err)
				}

				defer db.Close()

				Delete(db)

				for _, v := range news.Articles {
					Insert(db, v.Id, v.Title, v.Score, v.Time)
				}

			} else {
				fmt.Print("not")
				w.Write(HNmock())
			}

		}

	}

}
func GetDetails(a *Article, path string) {

	serverUrl := (path + "/item/" + strconv.Itoa(a.Id) + ".json?print=pretty")
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
func GetIDs(path string) [10]int {
	var IDs [10]int
	var body []byte

	urlA, err := url.Parse(path + "/topstories.json?print=pretty")
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
func CreateTable(db *sql.DB) {
	_, err := db.Exec(`Create table articles (id INT PRIMARY KEY, title STRING, score INT, time STRING)`)
	if err != nil {
		log.Fatal(err)
	}
}
func Insert(db *sql.DB, id int, title string, score int, time string) {
	_, err := db.Exec(`INSERT INTO articles(id, title, score, time) VALUES (?, ?, ?, ?)`, id, title, score, time)
	if err != nil {
		log.Fatal(err)
	}
}
func Delete(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM articles`)
	if err != nil {
		log.Fatal(err)
	}
}
func CheckDB() []Article {
	localArticals := make([]Article, 0, 10)

	nowAsString := time.Now().Format("2006-01-02 15:04:05")
	nowAsTime, err := time.Parse("2006-01-02 15:04:05", nowAsString)
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query(`SELECT * FROM articles `)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var article Article

		rows.Scan(&article.Id, &article.Title, &article.Score, &article.Time)

		articleTimeAsTime, err := time.Parse("2006-01-02 15:04:05", article.Time)
		if err != nil {
			fmt.Println(err)
		}

		diff := nowAsTime.Sub(articleTimeAsTime)
		hours := float32(diff.Hours())

		if hours <= 1 {
			localArticals = append(localArticals, article)
		}
	}

	return localArticals

}
func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/top", HandlerNewsMarshal("https://hacker-news.firebaseio.com/v0"))
	router.HandleFunc("/top", HandlerNewsHTMLTemplate("https://hacker-news.firebaseio.com/v0"))
	http.ListenAndServe(":9000", router)
}
