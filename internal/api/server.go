package api

import (
	"log"

	"aipi/internal/api/handlers"

	docs "aipi/docs"
	svc "aipi/internal/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router  *gin.Engine
	service *svc.Service
}

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	return &Server{
		router:  gin.Default(),
		service: svc.NewService(),
	}
}

// Register handlers
func (s *Server) SetUpRoutes() {
	r := s.router
	surl := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/", handlers.GetDocs)
	r.POST("/imgtoad", handlers.ImageHandler(s.service))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, surl))
}

func (s *Server) Run() {
	// Ideally set router pass by cfg
	log.Printf("Server running on port %d\n", 8080)
	s.router.Run(":8080")
}
