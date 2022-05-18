package handlers

import (
	"encoding/csv"
	"encoding/json"
	"final/cmd/echo/currentUser"
	"final/cmd/echo/helpers"
	db "final/cmd/echo/repository"
	"final/data"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jmoiron/jsonq"
	"github.com/labstack/echo"
)

type Weather struct {
	FormatedTemp string `json:"formatedTemp"`
	Description  string `json:"description"`
	City         string `json:"city"`
}

func GetAllLists(ctx echo.Context) error {
	myDB := db.GetDB()
	lists, err := myDB.GetLists(int32(currentUser.User.ID))
	if err != nil {
		log.Fatal(err)
	}

	if lists != nil {
		return ctx.JSON(http.StatusOK, lists)
	} else {
		return ctx.JSON(http.StatusOK, make([]string, 0))
	}
}
func PostList(ctx echo.Context) error {
	myDB := db.GetDB()

	var reqList data.List
	if err := ctx.Bind(&reqList); err != nil {
		return err
	}
	reqList.UserID = int32(currentUser.User.ID)
	myDB.InsertList(reqList)

	return ctx.JSON(http.StatusOK, reqList)
}
func DeleteList(ctx echo.Context) error {
	myDB := db.GetDB()
	id, _ := strconv.Atoi(ctx.Param("id"))
	myDB.DeleteList(id)

	return ctx.NoContent(http.StatusOK)
}
func GetTasks(ctx echo.Context) error {
	myDB := db.GetDB()
	tasks := make([]data.Task, 0)

	id, _ := strconv.Atoi(ctx.Param("id"))
	tasks = myDB.GetTasks(id)

	if tasks != nil {
		return ctx.JSON(http.StatusOK, tasks)
	} else {
		return ctx.JSON(http.StatusOK, make([]string, 0))
	}
}
func PostTask(ctx echo.Context) error {
	myDB := db.GetDB()
	var task data.Task

	if err := ctx.Bind(&task); err != nil {
		return err
	}
	log.Print(ctx.Param("id"))
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	task.ListID = int32(id)
	task.Completed = false

	myDB.CreateTask(task)

	return ctx.JSON(http.StatusOK, task)

}
func DeleteTask(ctx echo.Context) error {
	myDB := db.GetDB()

	id, _ := strconv.Atoi(ctx.Param("id"))
	myDB.DeleteTask(id)
	return ctx.NoContent(http.StatusOK)
}
func ToggleTask(ctx echo.Context) error {
	myDB := db.GetDB()

	var reqTask data.Task
	if err := ctx.Bind(&reqTask); err != nil {
		return err
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	task := myDB.ToggleTask(id, reqTask.Completed)

	return ctx.JSON(http.StatusOK, task)
}
func ExportToFile(ctx echo.Context) error {
	myDB := db.GetDB()

	lists, err := myDB.GetLists(int32(currentUser.User.ID))
	tasks := make([]string, 0)

	f, err := os.Create("test2.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)

	for _, list := range lists {
		tasks = myDB.GetTextsOfTasksInsideOfList(int32(list.ID))
		if err := w.Write(tasks); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	w.Flush()

	r := ctx.Response()
	r.Header().Set("Content-Type", "text/csv")
	r.Header().Set("Content-Length", "1000")
	r.Header().Set("Content-Disposition", "attachment; filename= test2.csv")

	return ctx.Attachment("test2.csv", "test2.csv")

}

func OpenWeatherMap(ctx echo.Context) error {

	apiKey := "bbec0e6c8e6f0dfc2fab86c0a724ea5c"

	lon := ctx.Request().Header.Get("lon")
	lat := ctx.Request().Header.Get("lat")
	weather := Weather{}
	data := map[string]interface{}{}

	path := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + apiKey
	req, _ := http.NewRequest("GET", path, nil)
	res, _ := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	weather.Description, err = jq.String("weather", "0", "description")
	weather.City, err = jq.String("name")
	temp, err := jq.Float("main", "temp")

	weather.FormatedTemp = helpers.ConverKelvinToCelsium(temp)

	fmt.Println("weather:", weather)

	return ctx.JSON(http.StatusOK, weather)

}
