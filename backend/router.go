package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/controller"
)

func initializeRoutes(serv *gin.Engine) {
	serv.POST("/goshort", controller.ShortenURLHashPost)
	serv.POST("/gocustom", controller.ShortenURLCustomPost)
	serv.POST("/golong", controller.OriginalURLPost)
	serv.GET("/:url", controller.ShortRedirect)
}
