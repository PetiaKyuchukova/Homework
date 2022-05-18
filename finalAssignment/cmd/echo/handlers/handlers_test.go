package handlers

import (
	"database/sql"
	"final/cmd/echo/currentUser"
	"final/cmd/echo/repository"
	"final/data"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var mySQL *sql.DB

//var mockDB = map[string]*data.User{ "[email protected]": &User{"Jon Snow", "[email protected]"}  }
func MockDatabase() {
	var err error
	mySQL, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	mySQL.Exec(`CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
	 	text varchar NOT NULL,
		completed boolean NOT NULL,
		list_id INT REFERENCES lists(id)
	   );`)

	mySQL.Exec(`CREATE TABLE lists (
		id  SERIAL PRIMARY KEY,
		name varchar,
		user_id INT REFERENCES users(id)
		);`)
	mySQL.Exec(`CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		username varchar,
		password varchar NOT NULL
		);`)
	repository.SetDB(mySQL)
	repository.GetDB().CreateTask(data.Task{ID: 0, Text: "string", Completed: false, ListID: 1})
	repository.GetDB().InsertList(data.List{ID: 1, Name: "string", UserID: 0})
}

var listJSON = `{"id":0,"name":"string","user_id":0}
`
var listOflistsJSON = `[{"id":1,"name":"string","user_id":0}]
`
var taskJSON = `{"id":0,"text":"string","completed":false,"list_id":1}
`
var toggeledTaskJSON = `{"id":0,"text":"string","completed":true,"list_id":1}
`
var completedJSON = `{"completed": true}
`
var tasksListJSON = `[{"id":0,"text":"string","completed":false,"list_id":1}]
`
var weatherJson = `{"formatedTemp":"33°C","description":"clear sky","city":"Al Kufrah"}
`

func TestCreateList(t *testing.T) {
	// Setup

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/lists", strings.NewReader(listJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, PostList(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, listJSON, rec.Body.String())
	}
}
func TestCreateTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/lists/:id", strings.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/lists/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, PostTask(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, taskJSON, rec.Body.String())
	}
}
func TestGetTasks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/lists/:id/tasks", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/lists/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, GetTasks(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, tasksListJSON, rec.Body.String())
	}

}
func TestGetAllLists(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/lists", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	currentUser.User.ID = 0
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, GetAllLists(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, listOflistsJSON, rec.Body.String())
	}

}
func TestToggleTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "http://localhost:3000/api/tasks/:id", strings.NewReader(completedJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/lists/:id/tasks")
	ctx.SetParamNames("id")
	ctx.SetParamValues("0")
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, ToggleTask(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, toggeledTaskJSON, rec.Body.String())
	}
}

func TestDeleteList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/lists/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/lists/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, DeleteList(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, http.NoBody, rec.Body.String())
	}

}
func TestDeleteTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/tasks/:id", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/tasks/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("0")
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, DeleteTask(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestOpenWeatherMap(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/weather", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.Request().Header.Set("lon", "23")
	ctx.Request().Header.Set("lat", "23")
	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, OpenWeatherMap(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, weatherJson, rec.Body.String())
	}

}

func TestExportToFile(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/list/export", http.NoBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := rec.Result()

	MockDatabase()

	defer res.Body.Close()

	if assert.NoError(t, ExportToFile(ctx)) {
		assert.Equal(t, 200, rec.Code)

	}

}
