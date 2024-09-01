package handlers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"

	"aipi/internal/models"

	svc "aipi/internal/service"

	"github.com/gin-gonic/gin"
)

// @Summary Root endpoint
// @Description Root endpoint that redirects to the Swagger documentation
// @Tags root
// @Success 301
// @Router / [get]
func GetDocs(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	// c.IndentedJSON(http.StatusOK, gin.H{
	// 	"message": "Hello World",
	// })
}

// ErrorResponse represents an error response
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}

// ImageHandler godoc
// @Summary Generate marketing suggestions
// @Description Generate marketing suggestions based on an uploaded image file or image URL
// @Tags marketing
// @Accept multipart/form-data
// @Produce json
// @Param image formData file false "Image file to analyze"
// @Param url formData string false "URL of the image to analyze"
// @Success 200 {object} models.MarketingSuggestions
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
		// Dont omit type to avoid go deleting imports(used by swagger)
		var res models.MarketingSuggestions = svc.CreateImgRequest(encImg)

		c.JSON(http.StatusOK, res)
	}
}
