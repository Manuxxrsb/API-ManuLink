package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBURL genera la URL de conexión a la base de datos PostgreSQL
func DBURL() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	log.Println("PGHOST:", os.Getenv("PGHOST"))
	log.Println("PGUSER:", os.Getenv("PGUSER"))
	log.Println("PGDNAME:", os.Getenv("PGNAME"))
	log.Println("PGPORT:", os.Getenv("PGPORT"))


	DBHost := os.Getenv("PGHost")
	DBUser := os.Getenv("PGUser")
	DBPassword := os.Getenv("PGPassword")
	DBPort := os.Getenv("PGPort")
	DBName := os.Getenv("PGName")

	
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
}