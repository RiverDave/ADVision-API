package api

import (
	"net/http"

	"aipi/internal/api/handlers"

	svc "aipi/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	service *svc.Service
}

func NewServer() *Server {
	return &Server{
		router:  gin.Default(),
		service: svc.NewService(),
	}
}

func (s *Server) SetUpRoutes() {
	s.router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	s.router.GET("/checkhealth", handlers.CheckHealth)
	s.router.POST("/upload", handlers.ImageHandler(s.service))
}

func (s *Server) Run() {
	// Ideally set router pass by cfg
	s.router.Run()
}
