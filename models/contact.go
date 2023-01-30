package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Contact struct {
	Id        uuid.UUID `json:"id" form:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      *string   `json:"name" form:"name" gorm:"not null;size:50;"`
	Phone     *string   `json:"phone" form:"phone" gorm:"not null;size:15;"`
	UserId    *string   `json:"userId" form:"userId" gorm:"foreignKey;"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

func (self Contact) TableName() string {
	return "contacts"
}
