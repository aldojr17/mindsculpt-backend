package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendResponse(c *gin.Context, isSuccess bool, data interface{}, message string, code int) {
	response := map[string]interface{}{
		"is_success": isSuccess,
		"message":    message,
	}

	if isSuccess {
		response["data"] = data
	} else {
		response["error"] = data
	}

	c.JSON(code, response)
}

func ResponseOK(c *gin.Context, data interface{}, message string) {
	sendResponse(c, true, data, message, http.StatusOK)
}

func ResponseBadRequest(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusBadRequest)
}

func ResponseUnauthorized(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusUnauthorized)
}

func ResponseNotFound(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusNotFound)
}

func ResponseUnsupportedMediaType(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusUnsupportedMediaType)
}

func ResponseInternalServerError(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusInternalServerError)
}
