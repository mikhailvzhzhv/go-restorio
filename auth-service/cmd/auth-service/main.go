package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/test", func(ctx *gin.Context) {
		var requestBody struct {
			Message string `json:"message"`
		}

		// Парсим JSON-тело запроса
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON format",
			})
			return
		}

		// Возвращаем ответ
		ctx.JSON(http.StatusOK, gin.H{
			"Posted message": requestBody.Message,
		})
	})

	router.Run(":8081")
}
