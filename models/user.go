package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email     string    `json:"email" binding:"required,email" gorm:"column:email;type:text;unique;index;size:100;not null"`
	Password  string    `json:"password" binding:"required,min=1,max=25" gorm:"column:password;type:text;not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (self User) TableName() string {
	return "users"
}
