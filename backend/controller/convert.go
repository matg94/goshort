package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
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
	convertRequest := URLConvertRequest{}

	json.Unmarshal([]byte(val), &convertRequest)

	if readError != nil || convertRequest.URL == "" {
		handleError(c, 400)
		return
	}

	c.JSON(200, gin.H{
		"url": repo.ShortenURL(convertRequest.URL),
	})
}

func OriginalURLPost(c *gin.Context) {
	body := c.Request.Body
	val, readError := ioutil.ReadAll(body)
	convertRequest := URLConvertRequest{}

	json.Unmarshal([]byte(val), &convertRequest)

	if readError != nil || convertRequest.URL == "" {
		handleError(c, 400)
		return
	}

	c.JSON(200, gin.H{
		"url": repo.OriginalURL(convertRequest.URL),
	})
}
