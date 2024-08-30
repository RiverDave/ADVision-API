package service

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type Service struct {
	openai *openai.Client
}

func NewService() *Service {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apikey := os.Getenv("OPEN_AI_API_KEY")
	client := initOpenAiWrapper(apikey)

	return &Service{
		openai: client,
	}
}

func initOpenAiWrapper(apikey string) *openai.Client {
	client := openai.NewClient(apikey)

	if client == nil {
		log.Fatal("Error loading .env file")
	}

	return client
}

func (s *Service) Client() *openai.Client {
	return s.openai
}

func (s *Service) ProcessImage() {
}
