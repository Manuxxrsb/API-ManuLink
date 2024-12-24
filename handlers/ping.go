package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping( db *gorm.DB ) gin.HandlerFunc {
		return func(c *gin.Context) {
			dbinfo := fmt.Sprintf("%v", db)
			c.JSON(http.StatusOK, gin.H{
				"mensaje": "pong",
				"bd": dbinfo,
			})
		}
}
