package lib

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	Database
}

func NewRepository(db Database) Repository {
	return Repository{
		Database: db,
	}
}

func (self Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle == nil {
		log.Fatal("Transaction database not found")
		return self
	}

	self.Database.DB = trxHandle
	return self
}
