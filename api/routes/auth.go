package routes

import (
	"contacts-go/api/controllers"
	"contacts-go/lib"
)

type AuthRoutes struct {
	handler        lib.RequestHandler
	authController controllers.AuthController
}

func (self AuthRoutes) Setup() {
	api := self.handler.Gin.Group("/api")
	{
		api.POST("/auth", self.authController.RegisterUser)
		api.POST("/auth/login", self.authController.LogIn)
	}
}

func NewAuthRoutes(
	handler lib.RequestHandler,
	authController controllers.AuthController,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		authController: authController,
	}
}
