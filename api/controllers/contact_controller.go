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

func (self ContactController) GetContacts(ctx *gin.Context) {
	paramId := ctx.Param("userId")
	id, err := uuid.FromString(paramId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contacts, err := self.contactService.FindByUser(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"contacts": contacts})
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

	ctx.JSON(http.StatusCreated, gin.H{"contact": contact})
}

func (self ContactController) UpdateContact(ctx *gin.Context) {
	updateContact := models.Contact{}

	if err := ctx.BindJSON(&updateContact); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.GetErrorMsgs(err))
		return
	}

	contact, err := self.contactService.Update(updateContact)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"contact": contact})
}

func (self ContactController) DeleteContact(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := uuid.FromString(paramId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = self.contactService.Delete(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "Contact deleted"})
}
