package config

import (
	"github.com/joho/godotenv"
	"github.com/slavajs/SimpleAPI/internal/schemas"
	"log"
	"os"
	"strconv"
)

var ProjectConfig schemas.Config

func createConfig() schemas.Config {
	return schemas.Config{
		DB: schemas.DatabaseConfig{
			User:     getStrEnv("DB_USER_NAME", ""),
			Password: getStrEnv("DB_PASSWORD", ""),
			Host:     getStrEnv("DB_HOST", ""),
			Port:     getStrEnv("DB_PORT", ""),
			DBName:   getStrEnv("DB_NAME", ""),
		},
		JWT: schemas.JWTConfig{
			TokenDuration: getIntEnv("JWT_DURATION", 1),
			SecretString:  getStrEnv("JWT_KEY", ""),
		},
	}
}

func getStrEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		intVal, _ := strconv.Atoi(value)
		return intVal
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
