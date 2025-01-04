package user

import (
	"api-manulink/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create_User(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var User models.User
		if err := Request.ShouldBindJSON(&User); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"Respuesta": "Error al recibir la información " + err.Error()})
			log.Println("Error al recibir la información: ", err)
			return
		}
		if err := db.Create(&User).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"Mensaje": "Error al crear usuario " + err.Error()})
			log.Println("Error al crear usuario: ", err)
			return
		}

		Request.JSON(http.StatusOK, User)
	}
}