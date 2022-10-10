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
		api.POST("/contact", self.contactController.CreateContact)
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
