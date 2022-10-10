package services

import (
	"contacts-go/models"
	"contacts-go/repositories"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ContactService struct {
	repository repositories.ContactRepository
}

func NewContactService(repository repositories.ContactRepository) ContactService {
	return ContactService{
		repository: repository,
	}
}

func (self ContactService) WithTrx(trxHandle *gorm.DB) ContactService {
	self.repository = self.repository.WithTrx(trxHandle)
	return self
}

func (self ContactService) FindOne(id uuid.UUID) (contact models.Contact, err error) {
	return contact, self.repository.Find(&contact, id).Error
}
