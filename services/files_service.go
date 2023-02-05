package services

import (
	"bytes"
	"contacts-go/lib"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilesService struct {
	bucket lib.FilesBucket
}

func NewFilesService(bucket lib.FilesBucket) FilesService {
	return FilesService{
		bucket: bucket,
	}
}

func (self FilesService) UploadFile(
	filename string,
	fileHeader *multipart.FileHeader,
) (objectId string, err error) {

	file, err := fileHeader.Open()
	if err != nil {
		return objectId, err
	}
	defer file.Close()

	id, err := self.bucket.UploadFromStream(filename, file)
	return id.String(), err
}

func (self FilesService) GetFile(id string) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return buf, err
	}

	if _, err := self.bucket.DownloadToStream(objectId, buf); err != nil {
		return buf, err
	}

	return buf, nil
}
