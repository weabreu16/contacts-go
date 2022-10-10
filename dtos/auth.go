package dtos

import "contacts-go/models"

type Auth struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}
