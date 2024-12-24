package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBURL genera la URL de conexión a la base de datos PostgreSQL
func DBURL() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file, falling back to environment variables")
	}
	DBHost := getEnv("DBHost")
	DBUser := getEnv("DBUser")
	DBPassword := getEnv("DBPassword")
	DBPort := getEnv("DBPort")
	DBName := getEnv("DBName")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
}

// getEnv busca una variable en los secrets o en las variables de entorno
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Error: %s environment variable not set", key)
	}
	return value
}