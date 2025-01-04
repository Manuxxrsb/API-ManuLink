package user

import (
	"api-manulink/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get_User(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		id := informacion.Param("id")
		var User models.User
		if err := db.First(&User, id).Error; err != nil {
			informacion.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrada"})
			log.Println("Usuario no encontrado",err)
			return
		}

		informacion.JSON(http.StatusOK, User)
		log.Println("Usuario encontrado",User)
	}
}

func GetallUser(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		var Users []models.User
		err := db.Find(&Users).Error
		if err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar todas las Users"})
			log.Println("Error al buscar todas las Usuarios",err)
			return
		}
		informacion.JSON(http.StatusOK, Users)
		log.Println("Encontradas todas las Usuarios")
	}
}