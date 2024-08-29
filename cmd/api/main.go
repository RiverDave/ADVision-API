package main

import (
	"aipi/internal/api"
)

func main() {
	server := api.NewServer()
	server.SetUpRoutes()
	server.Run()
}
