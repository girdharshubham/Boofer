package service

import (
	"boofer/api"
	"github.com/gin-gonic/gin"
)

func V1Routes(r *gin.Engine) {
	v1 := r.Group("/v1", gin.BasicAuth(gin.Accounts{
		"girdharshubham": "password",
	}))
	{
		v1.GET("/healthz", api.Healthz)
		v1.GET("/about", api.About)
		v1.GET("/", api.Home)
	}
}
