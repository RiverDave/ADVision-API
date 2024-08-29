package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Systems are doing okay",
	})
}
