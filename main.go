package main

import (
	"io"
	"net/http"
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

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	initLogger()
	initMiddleware(server)
	initRouteHandlers(server)

	server.Run(":8080")
}

func initLogger() {
	file, _ := os.Create("myLog.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func initMiddleware(server *gin.Engine) {
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
	)
}

func initRouteHandlers(server *gin.Engine) {
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video successfully saved"})
			}
		})
	}

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
}
