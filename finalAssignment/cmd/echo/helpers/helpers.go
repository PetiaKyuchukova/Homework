package helpers

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

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
