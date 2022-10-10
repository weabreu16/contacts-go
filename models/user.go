package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email     string    `json:"email" gorm:"column:email;type:text;unique;index;size:100"`
	Password  []byte    `json:"password" gorm:"column:password;type:bytea"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (self User) TableName() string {
	return "users"
}
