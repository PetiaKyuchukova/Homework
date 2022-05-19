package main

import (
	"database/sql"
	"final/cmd"
	"final/cmd/echo/customcontext"
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

var pathToOWApi = "https://api.openweathermap.org/data/2.5/weather?"

func main() {
	handlers.SetPathToOWApi(pathToOWApi)
	router := echo.New()

	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &customcontext.CustomContext{c}
			return next(cc)
		}
	})

	mySQL, err := sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}

	repository.SetDB(mySQL)
	myDB := repository.GetDB()

	router.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		cc := c.(*customcontext.CustomContext)

		if cc.GetUserName() == username && cc.GetUserPassword() == password {
			return true, nil
		} else {
			user := myDB.GetUser(username)
			checker := helpers.CheckPasswordHash(password, user.Password)

			cc.SetUserPassword(password)
			cc.SetUserName(username)
			cc.SetUserId(user.ID)

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
