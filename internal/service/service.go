package service

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
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

type outputSchema struct {
	Advertisement               string   `json:"advertisement"`
	CallToAction                string   `json:"call_to_action"`
	AltText                     string   `json:"alt_text"`
	ProductSuggestions          []string `json:"product_suggestions"`
	TargetAudienceInsights      string   `json:"target_audience_insights"`
	EmotionalToneAnalysis       string   `json:"emotional_tone_analysis"`
	SEOKeywords                 []string `json:"seo_keywords"`
	SocialMediaCaption          string   `json:"social_media_caption"`
	ContentIdeas                []string `json:"content_ideas"`
	Hashtags                    []string `json:"hashtags"`
	MarketingStrategyTips       string   `json:"marketing_strategy_tips"`
	ImageEnhancementSuggestions string   `json:"image_enhancement_suggestions"`
	CulturalAdaptations         string   `json:"cultural_adaptations"`
	LegalEthicalConsiderations  string   `json:"legal_ethical_considerations"`
	Emojis                      []string `json:"emojis"`
}

// Prompt user for image
// pass string previously encoded to base64
func (s *Service) CreateImgRequest(base64Image string, objective string) (outputSchema, error) {
	client := s.Client()

	// Generate obj to bind
	var out outputSchema
	schema, err := jsonschema.GenerateSchemaForType(out)
	if err != nil {
		log.Fatal("Failed to generate schema\n")
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT4oMini,
			MaxTokens: 500,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a marketing specialist working for a company that sells products online. You are tasked with creating a marketing advertisement based on the image provided. provide an altText for the image, a seoDescription, hashtags(And optional data for the other fields), And give a call to action (or call to actions) dont use any emojis in the advertisement field. The emojis field should contain only emojis that are relevant to the advertisement.",
				},
				{
					Role: openai.ChatMessageRoleUser,
					MultiContent: []openai.ChatMessagePart{
						{
							Type: openai.ChatMessagePartTypeText,
							Text: "Write a creative and engaging marketing advertisement for a product or service that aligns with the following image.",
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
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "advertisement",
					Schema: schema,
					Strict: true,
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("Error Processing chat request %v\n", err)
	}

	// parse json
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &out)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v\n", err)
	}

	return out, err
}
