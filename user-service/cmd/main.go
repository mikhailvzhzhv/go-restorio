package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/test", func(ctx *gin.Context) {
		msg := ctx.PostForm("message")
		ctx.JSON(http.StatusOK, gin.H{
			"Posted message": msg,
		})
	})

	router.Run(":8081")
}
