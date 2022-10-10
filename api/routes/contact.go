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
		api.GET("/contact/:id", self.contactController.GetContact)
		api.GET("/contact/user/:userId", self.contactController.GetContacts)
		api.POST("/contact", self.contactController.CreateContact)
		api.PUT("/contact", self.contactController.UpdateContact)
		api.DELETE("/contact/:id", self.contactController.DeleteContact)
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
