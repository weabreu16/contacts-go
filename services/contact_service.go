package services

import (
	"bytes"
	"contacts-go/lib"
	"contacts-go/models"
	"mime/multipart"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ContactService struct {
	repository   lib.Repository
	filesService FilesService
}

func NewContactService(repository lib.Repository, filesService FilesService) ContactService {
	return ContactService{
		repository:   repository,
		filesService: filesService,
	}
}

func (self ContactService) WithTrx(trxHandle *gorm.DB) ContactService {
	self.repository = self.repository.WithTrx(trxHandle)
	return self
}

func (self ContactService) Find(filters models.Contact) (contacts []models.Contact, err error) {
	return contacts, self.repository.Find(&contacts, &filters).Error
}

func (self ContactService) FindOne(id uuid.UUID) (contact models.Contact, err error) {
	return contact, self.repository.Find(&contact, id).Error
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

func (self ContactService) UploadImage(id uuid.UUID, fileHeader *multipart.FileHeader) (contact models.Contact, err error) {
	contact, err = self.FindOne(id)
	if err != nil {
		return contact, err
	}

	oldImageId := contact.ImageId

	imageId, err := self.filesService.UploadFile(fileHeader.Filename, fileHeader)
	if err != nil {
		return contact, err
	}

	if oldImageId != nil {
		self.filesService.RemoveFile(*oldImageId)
	}

	contact, err = self.Update(id, models.Contact{ImageId: &imageId})
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func (self ContactService) GetImage(id string) (*bytes.Buffer, error) {
	return self.filesService.GetFile(id)
}

func (self ContactService) RemoveImage(id string) error {
	return self.filesService.RemoveFile(id)
}
