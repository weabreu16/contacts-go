package repositories

import (
	"contacts-go/lib"
	"log"

	"gorm.io/gorm"
)

type ContactRepository struct {
	lib.Database
}

func NewContactRepository(db lib.Database) ContactRepository {
	return ContactRepository{
		Database: db,
	}
}

func (self ContactRepository) WithTrx(trxHandle *gorm.DB) ContactRepository {
	if trxHandle == nil {
		log.Fatal("Transaction database not found")
		return self
	}

	self.Database.DB = trxHandle
	return self
}
