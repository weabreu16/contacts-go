package services

import (
	"contacts-go/lib"
	"contacts-go/models"
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
	env lib.Env
}

func NewJWTService(env lib.Env) JWTService {
	return JWTService{
		env: env,
	}
}

func (self JWTService) Authorize(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(self.env.JWT_SECRET), nil
	})

	if token.Valid {
		return true, nil
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("Token malformed")
		}

		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New("Token expired")
		}
	}

	return false, errors.New("Couldn't handle token")
}

func (self JWTService) CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(self.env.JWT_SECRET))

	if err != nil {
		log.Fatal("JWT validation failed: ", err)
	}

	return tokenString
}
