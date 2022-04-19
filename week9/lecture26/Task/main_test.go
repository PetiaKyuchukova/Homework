package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"
)

var article = Article{
	Id:    0,
	Title: "test",
	Score: 0,
	Time:  "",
}

func MockDatabase() *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	return db

}
func TestCreateTable(t *testing.T) {
	db := MockDatabase()

	CreateTable(db)

	_, table_check := db.Query("SELECT * FROM articles")

	if table_check != nil {
		t.Error("Table not exist!")
	}
}
func TestInsert(t *testing.T) {
	var articleDB Article
	db := MockDatabase()

	CreateTable(db)
	Insert(db, article.Id, article.Title, article.Score, article.Time)

	row, err := db.Query(`SELECT * FROM articles WHERE id=?`, article.Id)
	for row.Next() {
		row.Scan(&articleDB.Id, &articleDB.Title, &articleDB.Score, &articleDB.Time)
		if err != nil {
			fmt.Println(err)
		}
	}

	if reflect.DeepEqual(article, articleDB) == false {
		t.Error("The result is not expected! Expected:", article, "Actual:", articleDB)
	}

}
func TestDelete(t *testing.T) {
	var counter int
	db := MockDatabase()

	CreateTable(db)
	Delete(db)
	row, err := db.Query(`SELECT * FROM articles`)
	for row.Next() {
		counter++
		if err != nil {
			fmt.Println(err)
		}
	}

	if counter != 0 {
		t.Error("The result is not expected! Expected:", 0, "Actual:", counter)
	}

}
