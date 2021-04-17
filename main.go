package main

import (
	"io"
	"os"

	"github.com/brianfajardo/gin-test/controller"
	"github.com/brianfajardo/gin-test/middleware"
	"github.com/brianfajardo/gin-test/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	initLogger()
	initMiddleware(server)
	initRouteHandlers(server)

	server.Run(":8080")
}

func initMiddleware(server *gin.Engine) {
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
	)
}

func initRouteHandlers(server *gin.Engine) {
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
}

func initLogger() {
	file, _ := os.Create("myLog.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
