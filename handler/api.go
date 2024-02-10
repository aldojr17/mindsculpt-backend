package handler

import (
	"mindsculpt/api"
	"mindsculpt/initialize"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
}

func NewAPIHandler(app *initialize.Application) *APIHandler {
	return &APIHandler{}
}

func (h *APIHandler) GetGenerationModels(c *gin.Context) {
	resp, err := api.GetGenerationModels()
	if err != nil {
		InternalServerError(c, err)
		return
	}

	SuccessResponse(c, resp, "Successfully get models")
}
