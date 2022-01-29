package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/controller"
)

func initializeRoutes(serv *gin.Engine) {
	serv.POST("/shorten", controller.ShortenURLPost)
}
