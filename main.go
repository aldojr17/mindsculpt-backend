package main

import (
	"mindsculpt/initialize"
	"mindsculpt/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := initialize.InitApp()

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Routes(r, app)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
