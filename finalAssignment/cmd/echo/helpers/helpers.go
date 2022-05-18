package helpers

import (
	"database/sql"
	"final/cmd/echo/currentUser"
	db "final/cmd/echo/repository"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(id int, username string, password string) {
	myDB := db.GetDB()
	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	err = myDB.CreateUser(id, username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}
func CheckAuth(username, password string, c echo.Context) (bool, error) {
	myDB := db.GetDB()
	currentUser.User = myDB.GetUser(username)
	checker := CheckPasswordHash(password, currentUser.User.Password)

	return checker, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ConverListIdtoSQL(id int) sql.NullInt32 {
	var listID sql.NullInt32
	if id != 0 {
		listID.Int32 = int32(id)
		listID.Valid = true
	}
	return listID
}
func ConverKelvinToCelsium(temp float64) string {
	return strconv.Itoa(int(temp-273.15)) + "°C"
}
