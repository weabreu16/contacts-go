package routes

import (
	"contacts-go/api/controllers"
	"contacts-go/api/middlewares"
	"contacts-go/lib"
)

type ContactRoutes struct {
	handler           lib.RequestHandler
	contactController controllers.ContactController
	jwtMiddleware     middlewares.JWTMiddleware
}

func (self ContactRoutes) Setup() {
	api := self.handler.Gin.Group("/api", self.jwtMiddleware.Handler())
	{
		api.GET("/contacts", self.contactController.GetContacts)
		api.GET("/contacts/:id", self.contactController.GetContact)
		api.POST("/contacts", self.contactController.CreateContact)
		api.GET("/contacts/image/:id", self.contactController.GetImage)
		api.POST("/contacts/:id/image", self.contactController.UploadImage)
		api.PUT("/contacts/:id", self.contactController.UpdateContact)
		api.DELETE("/contacts/:id", self.contactController.DeleteContact)
		api.DELETE("/contacts/image/:id", self.contactController.DeleteImage)
	}
}

func NewContactRoutes(
	handler lib.RequestHandler,
	contactController controllers.ContactController,
	jwtMiddleware middlewares.JWTMiddleware,
) ContactRoutes {
	return ContactRoutes{
		handler:           handler,
		contactController: contactController,
		jwtMiddleware:     jwtMiddleware,
	}
}
