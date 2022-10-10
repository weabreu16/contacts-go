package models

import "time"

type Contact struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt"`
}

func (self Contact) TableName() string {
	return "contacts"
}
