package api

import (
	"net/http"

	"aipi/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	// cfg *cfg
	// TODO: Add config
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) SetUpRoutes() {
	s.router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	s.router.GET("/checkhealth", handlers.CheckHealth)
}

func (s *Server) Run() {
	// Ideally set router pass by cfg
	s.router.Run()
}
