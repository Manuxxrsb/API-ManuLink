package config

import (
	"fmt"
	"log"
	"os"
)

// DBURL genera la URL de conexión a la base de datos PostgreSQL
func DBURL() string {

	log.Println("HOST:", os.Getenv("PGHost"))
	log.Println("USER:", os.Getenv("PGUser"))
	log.Println("NAME:", os.Getenv("PGName"))
	log.Println("PORT:", os.Getenv("PGPort"))


	DBHost := os.Getenv("PGHost")
	DBUser := os.Getenv("PGUser")
	DBPassword := os.Getenv("PGPassword")
	DBPort := os.Getenv("PGPort")
	DBName := os.Getenv("PGName")

	
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
}