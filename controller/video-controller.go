package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranavwadekar98/go_basic_gin_framework/entity"
	"github.com/pranavwadekar98/go_basic_gin_framework/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{service: service}
}

func (c controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c controller) ShowAll(ctx *gin.Context) {
	var videos = c.service.FindAll()
	data := gin.H{
		"title":  "My Videos",
		"videos": videos,
	}
	fmt.Println(data)
	ctx.HTML(http.StatusOK, "index.html", data)
}
