package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Contact struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      *string   `json:"name" gorm:"not null;size:50;"`
	Phone     *string   `json:"phone" gorm:"not null;size:15;"`
	UserId    *string   `json:"userId" gorm:"foreignKey;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (self Contact) TableName() string {
	return "contacts"
}
