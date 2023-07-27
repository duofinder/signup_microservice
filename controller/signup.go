package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tckthecreator/clean_arch_go/model"
)

type Controller struct {
	service model.Service
}

func NewController(r *gin.Engine, s model.Service) {
	controller := &Controller{
		service: s,
	}

	go r.POST("/signup", controller.Signup)
	// r.GET("/health")
}

func (con *Controller) Signup(ctx *gin.Context) {
	var auth model.Auth

	if err := ctx.ShouldBindJSON(auth); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := con.service.Signup(&auth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
