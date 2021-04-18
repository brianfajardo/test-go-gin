package controller

import (
	"github.com/brianfajardo/gin-test/entity"
	"github.com/brianfajardo/gin-test/service"
	"github.com/brianfajardo/gin-test/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-friendly", validators.ValidateFriendlyTitle)

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

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	controller.service.Save(video)

	return nil
}
