package domain

type APIGetModelsResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Version float32 `json:"version"`
	Type    string  `json:"type"`
}

type APIGenerateImageRequest struct {
	Width                int    `json:"width" binding:"required"`
	Height               int    `json:"height" binding:"required"`
	NegativePromptUnclip string `json:"negative_prompt_unclip" binding:"required"`
	Query                string `json:"query" binding:"required"`
}

type APIGenerateImageRawResponse struct {
	UUID   string `json:"uuid"`
	Status string `json:"status"`
}

type APIGetGenerationRawResponse struct {
	UUID     string   `json:"uuid"`
	Status   string   `json:"status"`
	Images   []string `json:"images"`
	Censored bool     `json:"censored"`
}

type APIGenerateImageResponse struct {
	UUID     string `json:"uuid"`
	Status   string `json:"status"`
	ImageUrl string `json:"image_url"`
	Censored bool   `json:"censored"`
}
