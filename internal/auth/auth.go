package auth

import (
	//"fmt"
	"github.com/dmishashkov/SimpleAPI/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hash), err
}

func CheckPassword(password, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil, err
}

func CheckUserExists(login string) bool { // TODO: this func
	database := db.GetDB()
	statement := `SELECT EXISTS(SELECT 1 FROM "usersAuthData" WHERE login = ($1))`
	var ans bool
	database.QueryRow(statement, login).Scan(&ans)
	return ans
}
