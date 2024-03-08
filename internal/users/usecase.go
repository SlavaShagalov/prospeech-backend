package users

import (
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
)

type Usecase interface {
	List() ([]models.User, error)
	Get(id int64) (models.User, error)
	GetByUsername(username string) (models.User, error)
	FullUpdate(params *FullUpdateParams) (models.User, error)
	PartialUpdate(params *PartialUpdateParams) (models.User, error)
	UpdateAvatar(id int64, imgData []byte, filename string) (*models.User, error)
	Delete(id int64) error
}
