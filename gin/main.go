package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	r.Run(":3001")
}
