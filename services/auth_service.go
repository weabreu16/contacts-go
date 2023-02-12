package services

import (
	"contacts-go/dtos"
	"contacts-go/lib"
	"contacts-go/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	env        lib.Env
	repository lib.Repository
	jwtService JWTService
}

func NewAuthService(env lib.Env, repository lib.Repository, jwtService JWTService) AuthService {
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
	err = self.repository.Find(&user, "email = ?", loginUser.Email).Error

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

func (self AuthService) hashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), self.env.SALT_ROUNDS)

	return string(result), err
}

func (self AuthService) comparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
