package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matg94/goshort/repo"
)

func ShortRedirect(c *gin.Context) {
	shortenedURL := c.Param("url")
	originalURL := repo.OriginalURL(shortenedURL)
	c.Redirect(302, originalURL)
}
