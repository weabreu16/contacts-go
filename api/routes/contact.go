package routes

import (
	"contacts-go/api/controllers"
	"contacts-go/lib"
)

type ContactRoutes struct {
	handler           lib.RequestHandler
	contactController controllers.ContactController
}

func (self ContactRoutes) Setup() {
	api := self.handler.Gin.Group("/api")
	{
		api.GET("/contacts", self.contactController.GetContacts)
		api.GET("/contacts/:id", self.contactController.GetContact)
		api.POST("/contacts", self.contactController.CreateContact)
		api.PUT("/contacts/:id", self.contactController.UpdateContact)
		api.DELETE("/contacts/:id", self.contactController.DeleteContact)
	}
}

func NewContactRoutes(
	handler lib.RequestHandler,
	contactController controllers.ContactController,
) ContactRoutes {
	return ContactRoutes{
		handler:           handler,
		contactController: contactController,
	}
}
