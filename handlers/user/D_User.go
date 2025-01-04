package user

import (
	"api-manulink/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete_User(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID Invalido"})
			log.Println("ID Invalido: ", err)
			return
		}

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			log.Println("Usuario no encontrado: ", err)
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo al eliminar usuario"})
			log.Println("Fallo al eliminar usuario: ", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
		log.Println("Usuario eliminado")
	}
}