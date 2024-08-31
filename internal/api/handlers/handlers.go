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

		// ext := filepath.Ext(file.Filename)
		// if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Only .jpg, .jpeg, and .png files are allowed"})
		// 	return
		// }
		//
		// // Generate a unique filename
		// filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		//
		// if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		// 	return
		// }

		// call client

		c.JSON(http.StatusOK, gin.H{
			"message": "Image uploaded successfully",
			"output":  res.Choices[0].Message,
		})
	}
}
