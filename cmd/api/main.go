package main

import (
	"aipi/internal/api"
)

// @title ADVision API Documentatio
// @version 1.0
// @description prototype for API documentation for image processing and marketing suggestions
// @BasePath /
func main() {
	server := api.NewServer()
	server.SetUpRoutes()
	server.Run()
}
