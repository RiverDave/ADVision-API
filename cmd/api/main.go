package main

import (
	"log"
	"os"

	"aipi/internal/api"
	c "aipi/internal/config"

	"github.com/joho/godotenv"
)

// @title ADVision API Documentation
// @version 1.0
// @description Prototype for API documentation for image processing and marketing suggestions
// @BasePath /
func main() {
	/*
		      Load environment variables from file (If we're developing locally)
				  When deploying on gcloud enviroment variables are passed through path(I think)
					Which means this will always fail. Thats why we don't error check this function if we're in prod
	*/

	err := godotenv.Load()
	env := os.Getenv("ENVIRONMENT")

	if env != "prod" && err != nil {
		log.Fatalf("%v\n", err.Error())
	}

	cfg := c.InitConfig(os.Getenv("OPEN_AI_API_KEY"), env)

	server := api.NewServer(cfg)
	server.SetUpRoutes()
	server.Run()
}
