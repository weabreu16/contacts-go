package controllers

import (
	"contacts-go/dtos"
	"contacts-go/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
}

func NewContactController() ContactController {
	return ContactController{}
}

func (self ContactController) GetContact(ctx *gin.Context) {
	paramId := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{"contact": paramId})
}

func (self ContactController) CreateContact(ctx *gin.Context) {
	createContact := dtos.CreateContactDto{}

	if err := ctx.BindJSON(&createContact); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.GetErrorMsgs(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"contact": createContact})
}
