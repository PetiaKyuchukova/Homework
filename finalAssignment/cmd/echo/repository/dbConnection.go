package repository

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0041129115"
	dbname   = "TaskManager"
	sslmode  = "disable"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=%s",
	host, port, user, password, dbname, sslmode)
