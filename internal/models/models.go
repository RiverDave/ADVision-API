package models

// MarketingSuggestions represents the response from the image processing service
// swagger:model
type MarketingSuggestions struct {
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
