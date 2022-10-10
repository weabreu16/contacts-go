package repositories

import (
	"contacts-go/lib"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	lib.Database
}

func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{
		Database: db,
	}
}

func (self UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		log.Fatal("Transaction database not found")
		return self
	}

	self.Database.DB = trxHandle
	return self
}
