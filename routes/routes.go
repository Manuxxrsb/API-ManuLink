package routes

import (
	"api-manulink/handlers"

	"github.com/gin-gonic/gin"
)

func ConfiguraRutas(router *gin.Engine) {

	router.GET("/ping",handlers.Ping())

}