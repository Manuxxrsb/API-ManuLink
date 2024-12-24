package main

import (
	"api-manulink/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	router := gin.Default()

	routes.ConfiguraRutas(router)

	router.Run(":8080")

}