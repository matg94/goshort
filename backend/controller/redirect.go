package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ShortRedirect(c *gin.Context) {
	shortnedURL := c.Param("url")
	fmt.Println(shortnedURL)
	c.Redirect(302, "http://google.com")
}
