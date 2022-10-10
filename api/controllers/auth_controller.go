package controllers

import (
	"contacts-go/dtos"
	"contacts-go/lib"
	"contacts-go/models"
	"contacts-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (self AuthController) RegisterUser(ctx *gin.Context) {
	createUser := dtos.CreateUserDto{}

	if err := ctx.BindJSON(&createUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.GetErrorMsgs(err))
		return
	}

	user := models.User{
		Email:    createUser.Email,
		Password: []byte(createUser.Password),
	}

	auth, err := self.authService.Register(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": auth.User, "token": auth.Token})
}

func (self AuthController) LogIn(ctx *gin.Context) {
	loginUser := dtos.LoginUserDto{}

	if err := ctx.BindJSON(&loginUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, lib.GetErrorMsgs(err))
		return
	}

	auth, err := self.authService.LogIn(loginUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": auth.User, "token": auth.Token})
}
