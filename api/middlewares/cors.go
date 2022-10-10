package middlewares

import (
	"contacts-go/lib"

	"github.com/gin-contrib/cors"
)

type CorsMiddleware struct {
	handler lib.RequestHandler
}

func NewCorsMiddleware(handler lib.RequestHandler, env lib.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
	}
}

func (self CorsMiddleware) Setup() {
	self.handler.Gin.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
	}))
}
