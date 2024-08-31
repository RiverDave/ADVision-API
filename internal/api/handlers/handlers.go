package handlers

import (
	"encoding/base64"
	"log"
	"net/http"

	svc "aipi/internal/service"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Systems are doing okay",
	})
}

func ImageHandler(svc *svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer file.Close()
		log.Printf("Uploaded file: %s\n", header.Filename)

		// read file and convert to buffer

		buf := make([]byte, header.Size)
		_, err = file.Read(buf)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read file %v", err)
			return
		}

		encImg := base64.StdEncoding.EncodeToString(buf)
		res, err := svc.CreateImgRequest(encImg)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read file %v", err)
			return
		}

		// call client

		c.JSON(http.StatusOK, res)
	}
}
