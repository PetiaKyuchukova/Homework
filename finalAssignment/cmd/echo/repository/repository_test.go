package repository

import (
	"context"
	"database/sql"
	"final/data"
	"log"
	"reflect"
	"testing"
)

var list = data.List{
	Name: "testList",
}
var task = data.Task{
	Text: "testTask",
}

func MockDatabase() {
	var err error
	mySQL, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	SetDB(mySQL)
}
func TestCreateUser(t *testing.T) {
	expectedUser := data.User{
		ID:       1,
		Username: "test",
		Password: "test",
	}
	MockDatabase()
	MockDB := GetDB()
	MockDB.CreateTableUsers()
	err := MockDB.CreateUser(1, "test", "test")
	if err != nil {
		log.Fatal(err)
	}
	dbUser := MockDB.GetUser("test")
	if err != nil {
		log.Fatal(err)
	}
	if reflect.DeepEqual(expectedUser, dbUser) == false {
		t.Error("Not expected result!", expectedUser, dbUser)
	}

}
func TestCreateTableUser(t *testing.T) {

	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableUsers()
	_, table_check := MockDB.Db.Query("SELECT * FROM users")
	if table_check != nil {
		t.Error("Table lists not exist!")
	}

}
func TestCreateTableLists(t *testing.T) {

	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableLists()
	_, table_check := MockDB.Db.Query("SELECT * FROM lists")
	if table_check != nil {
		t.Error("Table lists not exist!")
	}

}
func TestCreateTableTasks(t *testing.T) {

	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableTasks()
	_, table_check := MockDB.Db.Query("SELECT * FROM tasks")
	if table_check != nil {
		t.Error("Table tasks not exist!")
	}
}
func TestInsertTaks(t *testing.T) {

	MockDatabase()
	MockDB := GetDB()

	queries := data.New(MockDB.Db)

	MockDB.CreateTableTasks()
	MockDB.CreateTask(task)
	expextedID := 0

	_, err := queries.GetTask(context.Background(), int64(expextedID))
	if err != nil {
		t.Error("Task is not exist!")
	}

	MockDB.CreateTask(task)
	expextedID = 1
	_, err = queries.GetTask(context.Background(), int64(expextedID))
	if err != nil {
		t.Error("Task is not exist!")
	}

}
func TestInsertList(t *testing.T) {

	MockDatabase()
	MockDB := GetDB()

	queries := data.New(MockDB.Db)

	MockDB.CreateTableLists()
	MockDB.InsertList(list)
	expextedID := 1

	_, err := queries.GetList(context.Background(), int64(expextedID))
	if err != nil {
		t.Error("List is not exist!")
	}
}
func TestGetAllLists(t *testing.T) {

	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableLists()
	MockDB.InsertList(list)
	MockDB.InsertList(list)

}
func TestDeleteList(t *testing.T) {

	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableLists()
	MockDB.CreateTableTasks()
	MockDB.InsertList(list)

	MockDB.DeleteList(1)
	_, err := MockDB.GetLists(1)

	if err != nil {
		t.Error("Result is not expected!")
	}
}
func TestGetTasks(t *testing.T) {

	task.ListID = 1

	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableLists()
	MockDB.CreateTableTasks()
	MockDB.InsertList(list)
	MockDB.CreateTask(task)

	tasks := MockDB.GetTasks(1)

	if len(tasks) != 1 {
		t.Error("Result is not expected!")
	}
}
func TestToggleTask(t *testing.T) {
	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableTasks()
	MockDB.CreateTask(task)
	expectetResult := true
	tasksDB := MockDB.ToggleTask(0, true)

	if expectetResult != tasksDB.Completed {
		t.Error("Result is not expected! Expected:", expectetResult)
	}

}
func TestDeleteTask(t *testing.T) {
	MockDatabase()
	MockDB := GetDB()

	MockDB.CreateTableTasks()
	MockDB.CreateTask(task)
	MockDB.DeleteTask(0)

	queries := data.New(MockDB.Db)

	_, err := queries.GetTask(context.Background(), 0)
	if err == nil {
		t.Error("Not expected result!")
	}

}
