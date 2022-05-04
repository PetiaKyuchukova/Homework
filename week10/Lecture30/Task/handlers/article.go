package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
	"topstories/article"
)

type Storage interface {
	Insert(id int, title string, score int, time string)
	Delete()
	CheckDB() []article.Article
}

type News struct {
	PageTitle string            `json:"PageTitle"`
	Articles  []article.Article `json:"top_stories"`
}

func HandlerHN_HTMLTemplate(storage Storage) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("../templates/layout.html"))
	return func(w http.ResponseWriter, r *http.Request) {

		as := article.NewArticleService("https://hacker-news.firebaseio.com/v0")
		var wg sync.WaitGroup
		var m sync.Mutex
		var news News
		news.PageTitle = "Top 10 stories"
		news.Articles = storage.CheckDB()

		if len(news.Articles) == 10 {
			tmpl.Execute(w, news)
			log.Print("FromDB")
		} else {
			log.Print("else")

			news.Articles = make([]article.Article, 10, 10)

			ids := as.GetIDs()

			for i := 0; i < len(ids); i++ {
				news.Articles[i].Id = ids[i]
				news.Articles[i].Time = time.Now().Format("2006-01-02 15:04:05")
			}

			for i, v := range news.Articles {
				wg.Add(1)
				go func(i int, v article.Article) {
					m.Lock()
					as.GetDetails(&news.Articles[i])
					m.Unlock()
					wg.Done()
				}(i, v)
			}

			wg.Wait()

			tmpl.Execute(w, news)
			storage.Delete()
			for _, v := range news.Articles {
				storage.Insert(v.Id, v.Title, v.Score, v.Time)
			}
		}

	}
}
func HandlerHN_Marshal(storage Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		as := article.NewArticleService("https://hacker-news.firebaseio.com/v0")

		var wg sync.WaitGroup
		var m sync.Mutex
		var news News
		news.PageTitle = "Top 10 stories"

		news.Articles = storage.CheckDB()

		if len(news.Articles) == 10 {
			empData, err := json.MarshalIndent(news, "", "\t")
			if err != nil {
				fmt.Println(err)
			}

			w.Write(empData)

		} else {

			ids := as.GetIDs()

			news.Articles = make([]article.Article, 10, 10)

			for i := 0; i < len(ids); i++ {

				news.Articles[i].Id = ids[i]
				news.Articles[i].Time = time.Now().Format("2006-01-02 15:04:05")

			}

			for i, v := range news.Articles {
				wg.Add(1)
				go func(i int, v article.Article) {
					m.Lock()
					as.GetDetails(&news.Articles[i])
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

			storage.Delete()
			for _, v := range news.Articles {
				storage.Insert(v.Id, v.Title, v.Score, v.Time)
			}

		}

	}

}
