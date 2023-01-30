package dtos

import "contacts-go/models"

type CreateContactDto struct {
	Name   string `json:"name" binding:"required,min=1,max=50"`
	Phone  string `json:"phone" binding:"required,min=1,max=15"`
	UserId string `json:"userId" binding:"required"`
}

func (self *CreateContactDto) ToModel() models.Contact {
	return models.Contact{
		Name:   self.Name,
		Phone:  self.Phone,
		UserId: self.UserId,
	}
}

type UpdateContactDto struct {
	Name  string `json:"name" binding:"omitempty,min=1,max=50"`
	Phone string `json:"phone" binding:"omitempty,min=1,max=15"`
}

func (self *UpdateContactDto) ToModel() models.Contact {
	return models.Contact{
		Name:  self.Name,
		Phone: self.Phone,
	}
}
