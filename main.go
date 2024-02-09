package main

import (
	"mindsculpt/initialize"
	log "mindsculpt/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := initialize.InitApp()
	log.Infof("%+v", app)

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
