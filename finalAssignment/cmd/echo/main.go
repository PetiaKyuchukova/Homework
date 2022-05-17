package main

import (
	"final/cmd"
	"final/cmd/echo/currentUser"
	"final/cmd/echo/handlers"
	"final/cmd/echo/helpers"
	"final/cmd/echo/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

var err error

func main() {
	router := echo.New()
	repository.InitDB()

	router.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		myDB := repository.GetDB()
		currentUser.User = myDB.GetUser(username)
		checker := helpers.CheckPasswordHash(password, currentUser.User.Password)

		return checker, nil
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
