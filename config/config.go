package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ProjectConfig Config

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}
type Config struct {
	DB DatabaseConfig
}

func createConfig() Config {
	return Config{
		DB: DatabaseConfig{
			User:     getEnv("USER_NAME", ""),
			Password: getEnv("PASSWORD", ""),
			Host:     getEnv("HOST", ""),
			Port:     getEnv("PORT", ""),
			DBName:   getEnv("DB_NAME", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func loadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print("Error loading env variables")
		panic(err)
	}
	log.Print("[loadEnv] Successfully loaded env variables")
}

func init() {
	loadEnv()
	ProjectConfig = createConfig()
	log.Print("[loadEnv] Successfully created project config")
}
