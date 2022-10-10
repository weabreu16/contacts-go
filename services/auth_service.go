package services

import (
	"contacts-go/dtos"
	"contacts-go/lib"
	"contacts-go/models"
	"contacts-go/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	env        lib.Env
	repository repositories.UserRepository
	jwtService JWTService
}

func NewAuthService(env lib.Env, repository repositories.UserRepository, jwtService JWTService) AuthService {
	return AuthService{
		env:        env,
		repository: repository,
		jwtService: jwtService,
	}
}

func (self AuthService) WithTrx(trxHandle *gorm.DB) AuthService {
	self.repository = self.repository.WithTrx(trxHandle)
	return self
}

func (self AuthService) Register(user models.User) (auth dtos.Auth, err error) {
	user.Password, err = self.hashPassword(user.Password)

	if err != nil {
		return auth, err
	}

	result := self.repository.Create(&user)

	if result.Error != nil {
		return auth, result.Error
	}

	token := self.jwtService.CreateToken(user)
	auth = dtos.Auth{
		User:  user,
		Token: token,
	}

	return auth, nil
}

func (self AuthService) LogIn(loginUser dtos.LoginUserDto) (auth dtos.Auth, err error) {
	var user models.User
	err = self.repository.Find(&user, "email = ", loginUser.Email).Error

	if err != nil {
		return auth, err
	}

	if ok := self.comparePassword(user.Password, loginUser.Password); !ok {
		return auth, errors.New("Invalid password")
	}

	token := self.jwtService.CreateToken(user)
	auth = dtos.Auth{
		User:  user,
		Token: token,
	}

	return auth, nil
}

func (self AuthService) hashPassword(password []byte) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword(password, self.env.SALT_ROUNDS)

	return hashedPassword, err
}

func (self AuthService) comparePassword(hashedPassword []byte, passwordString string) bool {
	password := []byte(passwordString)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err == nil
}
