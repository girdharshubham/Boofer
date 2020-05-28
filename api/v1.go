package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(http.StatusOK, gin.H{
		"home": "Welcome back " + user,
	})
}

func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "All systems go",
	})
}

func About(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"about": "Hey there?",
	})
}
