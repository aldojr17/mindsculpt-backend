package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendResponse(c *gin.Context, isSuccess bool, data interface{}, message string, code int) {
	response := map[string]interface{}{
		"is_success": isSuccess,
		"data":       data,
		"message":    message,
	}
	c.JSON(code, response)
}

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	sendResponse(c, true, data, message, http.StatusOK)
}

func BadResponseError(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusBadRequest)
}

func InternalServerError(c *gin.Context, err error) {
	sendResponse(c, false, nil, err.Error(), http.StatusInternalServerError)
}
