package repository

import (
	"context"
	"database/sql"
	"final/cmd/echo/currentUser"
	"final/data"
	"log"

	_ "github.com/lib/pq"

	_ "modernc.org/sqlite"
)

type Database struct {
	Db *sql.DB
}

var mySQL *sql.DB

func GetDB() *Database {
	return &Database{Db: mySQL}
}

func InitDB() {
	var err error
	mySQL, err = sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}
}
func (repo *Database) GetUser(username string) data.User {
	var err error
	queries := data.New(repo.Db)
	currentUser.User, err = queries.GetUser(context.Background(), username)
	if err != nil {
		log.Fatal(err)
	}
	return currentUser.User

}
func (repo *Database) CreateUser(id int, username string, password string) error {

	queries := data.New(repo.Db)
	_, err := queries.CreateUser(context.Background(), data.CreateUserParams{ID: int64(id), Username: username, Password: password})
	if err != nil {
		return err
	}
	return nil
}
func (repo *Database) CreateTableLists() {
	_, err := mySQL.Exec(`CREATE TABLE lists (
		id  SERIAL PRIMARY KEY,
		name varchar,
		user_id INT REFERENCES users(id)
		);`)
	if err != nil {
		log.Fatal(err)
	}
}
func (repo *Database) CreateTableTasks() {
	_, err := mySQL.Exec(`CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
	 	text varchar NOT NULL,
		completed boolean NOT NULL,
		list_id INT REFERENCES lists(id)
	   );`)
	if err != nil {
		log.Fatal(err)
	}

}
func (repo *Database) CreateTableUsers() {
	_, err := mySQL.Exec(`CREATE TABLE users (
		id SERIAL PRIMARY KEY,
	 	username varchar,
	 	password varchar NOT NULL
	   );`)
	if err != nil {
		log.Fatal(err)
	}

}

func (repo *Database) SetListID_PK(queries *data.Queries) int32 {
	var listId int32

	id, err := queries.MaxIdlist(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if id == nil {
		listId = 1
	} else {
		listId = int32(id.(int64) + 1)
	}
	return listId

}
func (repo *Database) InsertList(list data.List) {
	queries := data.New(repo.Db)
	list.ID = int64(repo.SetListID_PK(queries))

	_, err := queries.CreateList(context.Background(), data.CreateListParams{ID: list.ID, Name: list.Name, UserID: list.UserID})
	if err != nil {
		log.Fatal(err)
	}

}
func (repo *Database) GetLists(id int32) ([]data.List, error) {
	queries := data.New(repo.Db)
	list, err := queries.GetUserLists(context.Background(), id)

	return list, err
}
func (repo *Database) DeleteList(id int) {
	queries := data.New(repo.Db)
	err := queries.DeleteList(context.Background(), int64(id))
	if err != nil {
		log.Fatal(err)
	}
	err = queries.DeleteTasksInsideList(context.Background(), int32(id))
	if err != nil {
		log.Fatal(err)
	}
}
func (repo *Database) SetTaskID_PK(queries *data.Queries) int32 {
	var taskId int32

	id, err := queries.MaxIdtask(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	if id == nil {
		taskId = 0
	} else {
		taskId = int32(id.(int64) + 1)
	}
	return taskId

}
func (repo *Database) GetTextsOfTasksInsideOfList(list_id int32) []string {
	queries := data.New(repo.Db)
	texts, err := queries.GetTextOfTasksInsideOfList(context.Background(), list_id)
	if err != nil {
		log.Fatal(err)
	}
	return texts

}
func (repo *Database) SetListID_FK(id int) sql.NullInt32 {
	var listID sql.NullInt32
	if id != 0 {
		listID.Int32 = int32(id)
		listID.Valid = true
	}
	return listID
}
func (repo *Database) CreateTask(task data.Task) {
	queries := data.New(repo.Db)
	task.ID = int64(repo.SetTaskID_PK(queries))

	_, err := queries.CreateTask(context.Background(), data.CreateTaskParams{ID: task.ID, Text: task.Text, ListID: task.ListID, Completed: false})
	if err != nil {
		log.Fatal(err)
	}

}
func (repo *Database) GetTasks(list_id int) []data.Task {
	queries := data.New(repo.Db)

	tasks, err := queries.GetTasksInsideOfList(context.Background(), int32(list_id))
	if err != nil {
		log.Fatal(err)
	}
	return tasks

}
func (repo *Database) ToggleTask(id int, completed bool) data.Task {
	queries := data.New(repo.Db)

	err := queries.ToggleTask(context.Background(), data.ToggleTaskParams{ID: int64(id), Completed: completed})
	if err != nil {
		log.Fatal(err)
	}
	task, err := queries.GetTask(context.Background(), int64(id))
	if err != nil {
		log.Fatal(err)
	}
	return task
}
func (repo *Database) DeleteTask(id int) {
	queries := data.New(repo.Db)
	err := queries.DeleteTask(context.Background(), int64(id))
	if err != nil {
		log.Fatal(err)
	}
}
