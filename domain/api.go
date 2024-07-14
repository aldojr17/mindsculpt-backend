package domain

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type APIGetModelsResponse struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Version float32 `json:"version"`
	Type    string  `json:"type"`
}

type APIGenerateImageRequest struct {
	Width                int    `json:"width"`
	Height               int    `json:"height"`
	NegativePromptUnclip string `json:"negative_prompt_unclip"`
	Query                string `json:"query" binding:"required"`
	ModelID              int    `json:"model_id"`
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

func (v *APIGenerateImageRequest) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return err
	}

	if len(v.Query) > 1000 {
		return fmt.Errorf("invalid length (maximum 1000)")
	}

	if v.Width == 0 {
		v.Width = 1024
	}

	if v.Height == 0 {
		v.Height = 1024
	}

	if v.ModelID == 0 {
		v.ModelID = 4
	}

	return nil
}
