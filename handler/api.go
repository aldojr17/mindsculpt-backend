package handler

import (
	"mindsculpt/config"
	"mindsculpt/domain"
	"mindsculpt/initialize"
	"mindsculpt/repository/cache"
	"mindsculpt/service"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	service *service.APIService
}

func NewAPIHandler(app *initialize.Application) *APIHandler {
	return &APIHandler{
		service: service.NewAPIService(cache.NewModelCache(app.Redis, config.GetConfig().Redis.GetTTLModel())),
	}
}

func (h *APIHandler) GetGenerationModels(c *gin.Context) {
	data, err := h.service.GetGenerationModels()
	if err != nil {
		InternalServerError(c, err)
		return
	}

	SuccessResponse(c, data, "Successfully get models")
}

func (h *APIHandler) GenerateImage(c *gin.Context) {
	var payload domain.APIGenerateImageRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		BadResponseError(c, err)
		return
	}

	data, err := h.service.GenerateImage(payload)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	SuccessResponse(c, data, "Successfully generate image")
}
