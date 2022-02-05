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
	server := gin.Default()
	server.Use(cors.Default())
	config.ReadConfig(fmt.Sprintf("config/%s.yaml", os.Args[1]))
	initializeRoutes(server)
	repo.InitRedis()
	server.Run()
}
