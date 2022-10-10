package dtos

type CreateContactDto struct {
	Name   string `json:"name" binding:"required,min=1,max=50"`
	Phone  string `json:"phone" binding:"required,min=1,max=15"`
	UserId string `json:"userId" binding:"required"`
}
