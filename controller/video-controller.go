package controller

import (
	"github.com/brianfajardo/gin-test/entity"
	"github.com/brianfajardo/gin-test/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service,
	}
}

func (controller *controller) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	controller.service.Save(video)

	return nil
}
