package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/config"
)

func main() {
	server := gin.Default()
	config.ReadConfig("config/dev.yaml")
	initializeRoutes(server)
	server.Run()
}
