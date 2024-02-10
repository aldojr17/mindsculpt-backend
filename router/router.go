package router

import (
	"mindsculpt/handler"
	"mindsculpt/initialize"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, app *initialize.Application) {
	apiHandler := handler.NewAPIHandler(app)
	router.GET("/get-models", apiHandler.GetGenerationModels)
}
