package handlers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"

	svc "aipi/internal/service"

	"github.com/gin-gonic/gin"
)

// @title API Documentation
// @version 1.0
// @description API documentation for image processing and marketing suggestions
// @BasePath /

// HelloWorld godoc
// @Summary Hello World endpoint
// @Description Returns a simple Hello World message
// @Tags example
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func HelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

// CheckHealth godoc
// @Summary Check Health endpoint
// @Description Check the health status of the system
// @Tags system
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func CheckHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Systems are doing okay",
	})
}

// MarketingSuggestions represents the response from the image processing service
// swagger:model
type MarketingSuggestions struct {
	// swagger:allOf
	// $ref: "./schemas/marketing_suggestions.yaml"
}

// ErrorResponse represents an error response
// swagger:model
type ErrorResponse struct {
	// swagger:allOf
	// $ref: "./schemas/error_response.yaml"
}

// ImageHandler godoc
// @Summary Generate marketing suggestions
// @Description Generate marketing suggestions based on an uploaded image file or image URL
// @Tags marketing
// @Accept multipart/form-data
// @Produce json
// @Param image formData file false "Image file to analyze"
// @Param url formData string false "URL of the image to analyze"
// @Success 200 {object} MarketingSuggestions
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /imgtoad [post]
func ImageHandler(svc *svc.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var imageData []byte
		var err error

		// Check if a file is uploaded
		file, _, err := c.Request.FormFile("image")
		if err == nil {
			// File was uploaded
			defer file.Close()
			imageData, err = io.ReadAll(file)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		} else {
			// Check if a URL is provided
			url := c.PostForm("url")
			if url != "" {
				// Fetch image from URL
				resp, err := http.Get(url)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}
				defer resp.Body.Close()

				imageData, err = io.ReadAll(resp.Body)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		// Log the size of the received image data
		log.Printf("Received image data of size: %d bytes\n", len(imageData))

		// Encode the image data to base64
		encImg := base64.StdEncoding.EncodeToString(imageData)

		// Process the image
		res, err := svc.CreateImgRequest(encImg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
