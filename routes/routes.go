package routes

import (
	"api-manulink/database"
	"api-manulink/handlers"
	"api-manulink/handlers/user"
	"api-manulink/models"
	"log"

	"github.com/gin-gonic/gin"
)

func ConfiguraRutas(router *gin.Engine) {

	db, err := database.OpenGormDB() //abro la conexion a la base de datos
	if err != nil {
		log.Fatalf("Error al conectarse a la BD: %v", err)
	}

	//Migracion de tablas a la bd
	db.AutoMigrate(&models.User{})
	


	router.GET("/ping",handlers.Ping(db))
	
	//Rutas de usuarios
	router.POST("/user", user.Create_User(db))
	router.GET("/user/:id", user.Get_User(db))
	router.GET("/users", user.GetallUser(db))
	router.DELETE("user/:id", user.Delete_User(db))

}