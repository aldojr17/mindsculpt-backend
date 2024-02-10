package domain

type APIGetModelsResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Version float32 `json:"version"`
	Type    string  `json:"type"`
}
