package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/slavajs/SimpleAPI/config"
	"log"
	_ "os"
)

func ConnectToDB(cfg config.DatabaseConfig) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("[ConnectToDB] error while connecting to DB")
	}
	log.Print("[ConnectToDB] successfully connected to DB")
	return db
}
