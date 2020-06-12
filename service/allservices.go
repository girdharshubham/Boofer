package service

import (
	"boofer/api"
	"github.com/gin-gonic/gin"
)

func V1Routes(r *gin.Engine) {
	{
		r.GET("/healthz", api.Healthz)
		r.GET("/about", api.About)
	}
	authenticated := r.Group("/v1", gin.BasicAuth(gin.Accounts{
		"girdharshubham": "password",
	}))
	{
		authenticated.GET("/", api.Home)
	}
}
