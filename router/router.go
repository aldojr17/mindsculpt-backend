package router

import (
	"mindsculpt/handler"
	"mindsculpt/initialize"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, app *initialize.Application) {
	apiHandler := handler.NewAPIHandler(app)
	router.GET("/get-models", apiHandler.GetGenerationModels)
	router.POST("/generate", apiHandler.GenerateImage)
	router.GET("/generation/:id", apiHandler.GetImageGeneration)
}
