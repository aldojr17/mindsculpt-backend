package handler

import (
	"mindsculpt/config"
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
