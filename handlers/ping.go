package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a simple handler to test the server status
// @Summary Ping the server
// @Description This will return a pong message
// @ID ping
// @Produce  json
// @Success 200 {string} string	"pong"
// @Router /ping [get]
// @Tags ping

func Ping() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		}
	}
