package chat

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	sync "sync"
	"time"
	"topstories/article"
	"topstories/repository"

	_ "modernc.org/sqlite"
)

type Server struct {
}
type News struct {
	PageTitle string            `json:"PageTitle"`
	Articles  []article.Article `json:"top_stories"`
}

func (s *Server) GetTopTenArticles(ctx context.Context, messRequest *MessageRequest) (*MessageReply, error) {
	mySQL, err := sql.Open("sqlite", "../data.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(mySQL)
	as := article.NewArticleService("https://hacker-news.firebaseio.com/v0")

	var wg sync.WaitGroup
	var m sync.Mutex
	var news News
	news.PageTitle = "Top 10 stories"

	news.Articles = repo.CheckDB()

	if len(news.Articles) == 10 {
		empData, err := json.MarshalIndent(news, "", "\t")
		if err != nil {
			fmt.Println(err)
		}

		return &MessageReply{Stories: string(empData)}, nil

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

		repo.Delete()
		for _, v := range news.Articles {
			repo.Insert(v.Id, v.Title, v.Score, v.Time)
		}
		return &MessageReply{Stories: string(empData)}, nil
	}

}

func (s *Server) mustEmbedUnimplementedHNServiceServer() {
}
