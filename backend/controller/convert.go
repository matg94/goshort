package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/models"
	"github.com/matg94/goshort/repo"
)

func handleError(c *gin.Context, code int) {
	c.JSON(code, gin.H{
		"error": "There was an error",
	})
}

func ShortenURLPost(c *gin.Context) {
	body := c.Request.Body
	val, readError := ioutil.ReadAll(body)
	shortenRequest := models.URL{}

	json.Unmarshal([]byte(val), &shortenRequest)

	if readError != nil || shortenRequest.Original == "" {
		handleError(c, 400)
		return
	}

	shortenRequest.Shortened = repo.ShortenURL(shortenRequest.Original)

	c.JSON(200, gin.H{
		"original":  shortenRequest.Original,
		"shortened": shortenRequest.Shortened,
	})
}
