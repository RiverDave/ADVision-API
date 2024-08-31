package service

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

// Wrap all the server Service providers if any, right here
type Service struct {
	openai *openai.Client
	// db... ?
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

// Prompt user for image
// pass string previously encoded to base64
func (s *Service) CreateImgRequest(base64Image string) (r openai.ChatCompletionResponse, err error) {
	client := s.Client()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an AI assistant capable of analyzing images.",
				},
				{
					Role: openai.ChatMessageRoleUser,
					MultiContent: []openai.ChatMessagePart{
						{
							Type: openai.ChatMessagePartTypeText,
							Text: "Analyze the following image:",
						},
						{
							Type: openai.ChatMessagePartTypeImageURL,
							ImageURL: &openai.ChatMessageImageURL{
								URL: "data:image/jpeg;base64," + base64Image,
							},
						},
					},
				},
			},
		},
	)
	return resp, err
}
