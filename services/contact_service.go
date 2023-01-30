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

func (self ContactService) FindByUser(userId uuid.UUID) (contacts []models.Contact, err error) {
	return contacts, self.repository.Find(&contacts, "userId = ?", userId).Error
}

func (self ContactService) Create(createContact models.Contact) (contact models.Contact, err error) {
	return createContact, self.repository.Create(&createContact).Error
}

func (self ContactService) Update(id uuid.UUID, updateContact models.Contact) (contact models.Contact, err error) {
	contact, err = self.FindOne(id)

	if err != nil {
		return contact, err
	}

	return contact, self.repository.Model(&contact).Updates(&updateContact).Error
}

func (self ContactService) Delete(id uuid.UUID) (err error) {
	return self.repository.Delete(models.Contact{}, id).Error
}
