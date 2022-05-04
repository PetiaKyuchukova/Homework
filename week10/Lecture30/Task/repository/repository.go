package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"topstories/article"
	"topstories/db"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (rp *Repository) Insert(id int, title string, score int, time string) {

	queries := db.New(rp.db)
	queries.CreateArticle(context.Background(), db.CreateArticleParams{
		ID:    int32(id),
		Title: title,
		Score: int32(score),
		Time:  time,
	})
}
func (rp *Repository) Delete() {

	queries := db.New(rp.db)
	queries.DeleteArticles(context.Background())

}
func (rp *Repository) CheckDB() []article.Article {

	queries := db.New(rp.db)

	localArticals := make([]article.Article, 0, 10)

	availableArticals, err := queries.GetArticles(context.Background())

	nowAsString := time.Now().Format("2006-01-02 15:04:05")
	nowAsTime, err := time.Parse("2006-01-02 15:04:05", nowAsString)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range availableArticals {
		articleTimeAsTime, err := time.Parse("2006-01-02 15:04:05", v.Time)
		if err != nil {
			fmt.Println(err)
		}

		diff := nowAsTime.Sub(articleTimeAsTime)
		hours := float32(diff.Hours())

		if hours <= 1 {
			article := article.Article{Id: int(v.ID), Title: v.Title, Score: int(v.Score), Time: v.Time}
			localArticals = append(localArticals, article)
		}

	}

	return localArticals

}
