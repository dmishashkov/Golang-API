package auth

import (
	"fmt"
	"github.com/slavajs/SimpleAPI/config"
	"github.com/slavajs/SimpleAPI/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

func CheckUserExists(login string) { // TODO: this gfunc
	database := db.ConnectToDB(config.ProjectConfig.DB)
}
