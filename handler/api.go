package handler

import (
	"mindsculpt/config"
	"mindsculpt/domain"
	"mindsculpt/initialize"
	"mindsculpt/repository"
	"mindsculpt/repository/cache"
	"mindsculpt/service"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	service *service.APIService
}

func NewAPIHandler(app *initialize.Application) *APIHandler {
	return &APIHandler{
		service: service.NewAPIService(
			cache.NewModelCache(app.Redis, config.GetConfig().Redis.GetTTLModel()),
			repository.NewImageGenerationRepository(app.Database),
		),
	}
}

func (h *APIHandler) GetGenerationModels(c *gin.Context) {
	data, err := h.service.GetGenerationModels()
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseOK(c, data, "Successfully get models")
}

func (h *APIHandler) GenerateImage(c *gin.Context) {
	var payload domain.APIGenerateImageRequest

	if err := payload.Validate(c); err != nil {
		ResponseBadRequest(c, err)
		return
	}

	data, err := h.service.GenerateImage(payload)
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseCreated(c, data, "Successfully generate image")
}

func (h *APIHandler) GetImageGeneration(c *gin.Context) {
	uuid := c.Param("id")

	data, err := h.service.GetImageGeneration(uuid)
	if err != nil {
		ResponseInternalServerError(c, err)
		return
	}

	ResponseOK(c, data, "Successfully get data")
}
