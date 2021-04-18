package controller

import (
	"net/http"
	"net/url"

	"github.com/brianfajardo/gin-test/entity"
	"github.com/brianfajardo/gin-test/service"
	"github.com/brianfajardo/gin-test/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	initValidators()

	return &controller{
		service,
	}
}

func initValidators() {
	validate = validator.New()
	validate.RegisterValidation("containsProfanity", validators.ValidateProfanity)
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	// Rewrite Youtube video urls if they are not embedded
	m, _ := url.Parse(video.Url)
	videoId := m.Query().Get("v")
	video.Url = "https://www.youtube.com/embed/" + videoId

	c.service.Save(video)

	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	templateData := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", templateData)
}
