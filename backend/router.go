package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/controller"
)

func initializeRoutes(serv *gin.Engine) {
	serv.POST("/short", controller.ShortenURLHashPost)
	serv.POST("/custom", controller.ShortenURLCustomPost)
	serv.POST("/long", controller.OriginalURLPost)
	serv.GET("/:url", controller.ShortRedirect)
}
