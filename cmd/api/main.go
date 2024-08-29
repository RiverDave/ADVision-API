package main

import (
	"aipi/internal/api"
)

func main() {
	server := api.NewServer()
	server.Run()
}
