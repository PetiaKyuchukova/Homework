package main

import (
	"database/sql"
	"final/cmd"
	"final/cmd/echo/currentUser"
	"final/cmd/echo/handlers"
	"final/cmd/echo/helpers"
	"final/cmd/echo/repository"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

var err error
var isLogg bool = false

func main() {
	router := echo.New()

	mySQL, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}

	repository.SetDB(mySQL)

	router.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if isLogg == true {
			return true, nil
		} else {
			myDB := repository.GetDB()
			currentUser.User = myDB.GetUser(username)
			checker := helpers.CheckPasswordHash(password, currentUser.User.Password)
			log.Print("A", checker)
			log.Print("middleware")
			if checker == true {
				isLogg = true
			} else {
				isLogg = false
			}
			return checker, nil
		}
	}))

	//Lists API endpoints
	router.GET("/api/lists", handlers.GetAllLists)
	router.POST("/api/lists", handlers.PostList)
	router.DELETE("/api/lists/:id", handlers.DeleteList)

	//Tasks API endpoints
	router.GET("api/lists/:id/tasks", handlers.GetTasks)
	router.POST("/api/lists/:id/tasks", handlers.PostTask)
	router.DELETE("api/tasks/:id", handlers.DeleteTask)
	router.PATCH("api/tasks/:id", handlers.ToggleTask)

	//Export data to file API endpoint
	router.GET("api/list/export", handlers.ExportToFile)

	//OpenWeatherMap API endpoint
	router.GET("api/weather", handlers.OpenWeatherMap)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))

}
