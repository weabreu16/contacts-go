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

func (self ContactController) GetContacts(ctx *gin.Context) {
	filters := models.Contact{}

	if err := ctx.BindQuery(&filters); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	contacts, err := self.contactService.Find(filters)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contacts)
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

	ctx.JSON(http.StatusOK, contact)
}

func (self ContactController) CreateContact(ctx *gin.Context) {
	createContact := dtos.CreateContactDto{}

	if err := ctx.BindJSON(&createContact); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.GetErrorMsgs(err))
		return
	}

	contact, err := self.contactService.Create(createContact.ToModel())

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, contact)
}

func (self ContactController) UpdateContact(ctx *gin.Context) {
	paramId := ctx.Param("id")
	updateContact := dtos.UpdateContactDto{}

	id, err := uuid.FromString(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.BindJSON(&updateContact); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.GetErrorMsgs(err))
		return
	}

	contact, err := self.contactService.Update(id, updateContact.ToModel())

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, contact)
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

func (self ContactController) GetImage(ctx *gin.Context) {
	paramId := ctx.Param("id")

	image, err := self.contactService.GetImage(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error(), "image": image})
		return
	}

	ctx.Writer.Write(image.Bytes())
}

func (self ContactController) UploadImage(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, err := uuid.FromString(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := self.contactService.UploadImage(id, fileHeader)

	ctx.JSON(http.StatusOK, contact)
}
