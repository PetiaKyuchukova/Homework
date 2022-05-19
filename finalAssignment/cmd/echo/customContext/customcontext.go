package customcontext

import "github.com/labstack/echo"

var userID int64
var isLogg bool
var userName string
var pass string

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) SetUserId(id int64) {
	userID = id
}

func (c *CustomContext) GetUserId() int64 {
	return userID
}
func (c *CustomContext) SetUserName(username string) {
	userName = username
}

func (c *CustomContext) GetUserName() string {
	return userName
}
func (c *CustomContext) SetUserPassword(password string) {
	pass = password
}

func (c *CustomContext) GetUserPassword() string {
	return pass
}
func (c *CustomContext) SetIsLogg(islogged bool) {
	isLogg = islogged
}

func (c *CustomContext) GetIsLogg() bool {
	return isLogg
}
