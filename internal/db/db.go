package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/slavajs/SimpleAPI/config"
	"github.com/slavajs/SimpleAPI/internal/schemas"
	"log"
	"sync"
)

func ConnectToDB(cfg schemas.DatabaseConfig) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("[ConnectToDB] error while connecting to DB", err)
	}
	log.Print("[ConnectToDB] successfully connected to DB")
	return db
}

var singleton sync.Once
var myDB *sql.DB

func GetDB() *sql.DB {
	singleton.Do(func() {
		myDB = ConnectToDB(config.ProjectConfig.DB)

	})
	return myDB
}
