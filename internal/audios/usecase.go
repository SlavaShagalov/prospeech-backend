package boards

import (
	"context"
	pFiles "github.com/SlavaShagalov/prospeech-backend/internal/files"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
)

type CreateParams struct {
	UserID int64
	File   pFiles.File
}

type Usecase interface {
	Create(ctx context.Context, params *CreateParams) (*models.Audio, error)
	List(ctx context.Context, userID int64) ([]models.Audio, error)
	//Get(ctx context.Context, id int64) (*models.Audio, error)
	//PartialUpdate(ctx context.Context, params *PartialUpdateParams) (*models.Audio, error)
	//Delete(ctx context.Context, id int64) error
}
