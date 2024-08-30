package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	svc "aipi/internal/service"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Systems are doing okay",
	})
}

func ProcessImage(svc *svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only .jpg, .jpeg, and .png files are allowed"})
			return
		}

		// Generate a unique filename
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

		if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// call client

		c.JSON(http.StatusOK, gin.H{
			"message":  "Image uploaded successfully",
			"filename": filename,
		})
	}
}
