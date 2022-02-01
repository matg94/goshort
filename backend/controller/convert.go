package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/repo"
)

func handleError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func ShortenURLHashPost(c *gin.Context) {
	body := c.Request.Body
	val, readError := ioutil.ReadAll(body)
	convertRequest := URLConvertRequest{}

	json.Unmarshal([]byte(val), &convertRequest)

	if readError != nil || convertRequest.URL == "" {
		handleError(c, 400, readError)
		return
	}

	c.JSON(200, gin.H{
		"url": repo.ShortenURLHash(convertRequest.URL),
	})
}

func ShortenURLCustomPost(c *gin.Context) {
	body := c.Request.Body
	val, readError := ioutil.ReadAll(body)
	customRequest := URLCustomConvertRequest{}

	json.Unmarshal([]byte(val), &customRequest)

	if readError != nil || customRequest.URL == "" || customRequest.Short == "" {
		handleError(c, 400, readError)
		return
	}

	var isValid = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if !isValid(customRequest.Short) {
		handleError(c, 400, fmt.Errorf("requested url does not meet requirements"))
	}

	short, err := repo.ShortenURLCustom(customRequest.URL, customRequest.Short)

	if err != nil {
		handleError(c, 409, err)
	}

	c.JSON(200, gin.H{
		"url": short,
	})
}

func OriginalURLPost(c *gin.Context) {
	body := c.Request.Body
	val, readError := ioutil.ReadAll(body)
	convertRequest := URLConvertRequest{}

	json.Unmarshal([]byte(val), &convertRequest)

	if readError != nil || convertRequest.URL == "" {
		handleError(c, 400, readError)
		return
	}

	c.JSON(200, gin.H{
		"url": repo.OriginalURL(convertRequest.URL),
	})
}
