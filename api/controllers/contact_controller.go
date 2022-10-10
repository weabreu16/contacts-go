package controllers

import (
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
