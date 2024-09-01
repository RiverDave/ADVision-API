package main

import (
	"log"

	"aipi/internal/api"

	"github.com/joho/godotenv"
)

// @title ADVision API Documentation
// @version 1.0
// @description Prototype for API documentation for image processing and marketing suggestions
// @BasePath /
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	server := api.NewServer()
	server.SetUpRoutes()
	server.Run()
}
