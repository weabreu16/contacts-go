package middlewares

import (
	"contacts-go/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTMiddleware struct {
	jwtService services.JWTService
}

func NewJWTMiddleware(jwtService services.JWTService) JWTMiddleware {
	return JWTMiddleware{
		jwtService: jwtService,
	}
}

func (self JWTMiddleware) Setup() {}

func (self JWTMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")

		if len(token) == 2 {
			authToken := token[1]
			authorized, err := self.jwtService.Authorize(authToken)

			if authorized {
				ctx.Next()
				return
			}

			log.Fatal(err)
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		ctx.Abort()
	}
}
