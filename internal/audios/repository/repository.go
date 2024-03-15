package repository

import (
	"context"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
	"time"
)

type CreateParams struct {
	UserID   int64
	Title    string
	URL      string
	Duration time.Duration
}

type PartialUpdateParams struct {
	ID                int
	Title             string
	UpdateTitle       bool
	Description       string
	UpdateDescription bool
	WorkspaceID       int
	UpdateWorkspaceID bool
}

type Repository interface {
	Create(ctx context.Context, params *CreateParams) (*models.Audio, error)
	List(ctx context.Context, userID int64) ([]models.Audio, error)
	Get(ctx context.Context, id int64) (*models.Audio, error)
	//PartialUpdate(ctx context.Context, params *PartialUpdateParams) (models.Audio, error)
	Delete(ctx context.Context, id int64) error
}
