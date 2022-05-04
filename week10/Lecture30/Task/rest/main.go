package main

import (
	"database/sql"
	"log"
	"net/http"
	handlers "topstories/handlers"
	repository "topstories/repository"

	_ "modernc.org/sqlite"
)

func main() {
	router := http.NewServeMux()
	mySQL, err := sql.Open("sqlite", "../data.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(mySQL)
	router.HandleFunc("/api/top", handlers.HandlerHN_Marshal(repo))
	router.HandleFunc("/top", handlers.HandlerHN_HTMLTemplate(repo))
	http.ListenAndServe(":9000", router)
}
