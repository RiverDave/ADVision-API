package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	// TODO: Add config
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) setUpRoutes() {
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}

func (s *Server) Run() {
	s.router.Run()
}
