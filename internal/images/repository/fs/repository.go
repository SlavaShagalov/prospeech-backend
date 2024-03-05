package fs

import (
	pImages "github.com/SlavaShagalov/prospeech-backend/internal/images"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"

	"go.uber.org/zap"
)

type repository struct {
	uploader *manager.Uploader
	log      *zap.Logger
}

func New(log *zap.Logger) pImages.Repository {
	return &repository{
		log: log,
	}
}

func (repo *repository) Create(imgName string, imgData []byte) (location string, err error) {
	repo.log.Debug("Image created")
	return "", nil
}

func (repo *repository) Get(location string) (imgData []byte, err error) {
	return nil, nil
}

func (repo *repository) Update(location string, imgData []byte) (err error) {
	repo.log.Debug("Image updated")
	return nil
}

func (repo *repository) Delete(location string) (err error) {
	return nil
}
