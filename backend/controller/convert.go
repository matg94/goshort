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

	shortenedURL, err := repo.ShortenURLHash(convertRequest.URL)

	if err != nil {
		handleError(c, 500, err)
		return
	}

	c.JSON(200, gin.H{
		"url": shortenedURL,
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
		return
	}

	short, err := repo.ShortenURLCustom(customRequest.URL, customRequest.Short)

	if err != nil {
		if err.Error() == "URL is already in use" {
			handleError(c, 409, err)
			return
		}
		handleError(c, 500, err)
		return
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

	originalURL := repo.OriginalURL(convertRequest.URL)

	if originalURL == "" {
		handleError(c, 404, fmt.Errorf("URL not found"))
	}

	c.JSON(200, gin.H{
		"url": repo.OriginalURL(convertRequest.URL),
	})
}
