package domain

type APIGetModelsResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Version float32 `json:"version"`
	Type    string  `json:"type"`
}

type APIGenerateImageRequest struct {
	UUID string `json:"uuid" binding:"required"`
}

type APIGenerateImageRawResponse struct {
	UUID     string   `json:"uuid"`
	Status   string   `json:"status"`
	Images   []string `json:"images"`
	Censored bool     `json:"censored"`
}

type APIGenerateImageResponse struct {
	UUID     string `json:"uuid"`
	Status   string `json:"status"`
	Image    string `json:"image"`
	Censored bool   `json:"censored"`
}
