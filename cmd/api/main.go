package main

import (
	"aipi/internal/api"
)

// @title ADVision API Documentation
// @version 1.0
// @description Prototype for API documentation for image processing and marketing suggestions
// @BasePath /
func main() {
	server := api.NewServer()
	server.SetUpRoutes()
	server.Run()
}
