package users

import (
	"context"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
)

type CreateParams struct {
	Name           string
	Username       string
	Email          string
	HashedPassword string
}

type FullUpdateParams struct {
	ID       int64
	Username string
	Email    string
	Name     string
}

type PartialUpdateParams struct {
	ID             int64
	Username       string
	UpdateUsername bool
	Email          string
	UpdateEmail    bool
	Name           string
	UpdateName     bool
}

type Repository interface {
	Create(ctx context.Context, params *CreateParams) (models.User, error)

	List(ctx context.Context) ([]models.User, error)
	Get(ctx context.Context, id int64) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)

	FullUpdate(ctx context.Context, params *FullUpdateParams) (models.User, error)
	PartialUpdate(ctx context.Context, params *PartialUpdateParams) (models.User, error)
	UpdateAvatar(ctx context.Context, id int64, avatar string) error

	Delete(ctx context.Context, id int64) error

	Exists(ctx context.Context, id int64) (bool, error)
}
