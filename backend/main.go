package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/config"
	"github.com/matg94/goshort/repo"
)

func main() {

	profile := os.Args[1]

	server := gin.Default()
	server.Use(cors.Default())
	if profile == "prod" {
		config.DownloadConfig("goshort/config/prod.yaml", "config/prod.yaml")
		fmt.Println("Config files downloaded.")
	}
	config.ReadConfig(fmt.Sprintf("config/%s.yaml", profile))
	initializeRoutes(server)
	repo.InitRedis()
	server.Run()
}
