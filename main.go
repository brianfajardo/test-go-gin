package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/api", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "ok"})
	})

	server.Run(":8080")
}
