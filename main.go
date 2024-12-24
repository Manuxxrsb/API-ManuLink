package main

import (
	"api-manulink/routes"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Iniciando API")
	router := gin.Default()

	//coors MIDLEWARE DE GIN
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))


	routes.ConfiguraRutas(router)

	log.Println("Conexión exitosa a la base de datos y API servida")

	router.Run(":8080")

	fmt.Println("API cerrada")

}
