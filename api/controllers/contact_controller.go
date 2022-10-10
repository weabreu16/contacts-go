package controllers

import (
	"contacts-go/dtos"
	"contacts-go/lib"
	"contacts-go/models"
	"contacts-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ContactController struct {
	contactService services.ContactService
}

func NewContactController(contactService services.ContactService) ContactController {
	return ContactController{
		contactService: contactService,
	}
}

func (self ContactController) GetContact(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := uuid.FromString(paramId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := self.contactService.FindOne(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"contact": contact})
}

func (self ContactController) CreateContact(ctx *gin.Context) {
	createContact := dtos.CreateContactDto{}

	if err := ctx.BindJSON(&createContact); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.GetErrorMsgs(err))
		return
	}

	contact := models.Contact{
		Name:   createContact.Name,
		Phone:  createContact.Phone,
		UserId: createContact.UserId,
	}

	contact, err := self.contactService.Create(contact)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"contact": contact})
}
