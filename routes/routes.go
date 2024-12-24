package routes

import (
	"api-manulink/database"
	"api-manulink/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func ConfiguraRutas(router *gin.Engine) {

	db, err := database.OpenGormDB() //abro la conexion a la base de datos
	if err != nil {
		log.Fatalf("Error al conectarse a la BD: %v", err)
	}


	router.GET("/ping",handlers.Ping(db))

}