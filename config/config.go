package config

import (
	"fmt"
	"os"
)

// DBURL genera la URL de conexión a la base de datos PostgreSQL
func DBURL() string {

	DBHost := os.Getenv("PGHost")
	DBUser := os.Getenv("PGUser")
	DBPassword := os.Getenv("PGPassword")
	DBPort := os.Getenv("PGPort")
	DBName := os.Getenv("PGName")

	
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DBUser, DBPassword, DBHost, DBPort, DBName)
}