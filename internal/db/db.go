package db

import (
	"database/sql"
	"fmt"
	"github.com/dmishashkov/SimpleAPI/config"
	"github.com/dmishashkov/SimpleAPI/internal/schemas"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func ConnectToDB(cfg schemas.DatabaseConfig) *sql.DB {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("[ConnectToDB] error while connecting to DB", err)
	}
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
