package main

import (
	"boofer/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	service.V1Routes(r)
	if ok := r.Run(":80"); ok != nil {
		panic("Something went wrong!")
	}

}
