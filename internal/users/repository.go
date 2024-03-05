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
	ID       int
	Username string
	Email    string
	Name     string
}

type PartialUpdateParams struct {
	ID             int
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
	Get(ctx context.Context, id int) (models.User, error)
	GetByUsername(ctx context.Context, username string) (models.User, error)

	FullUpdate(ctx context.Context, params *FullUpdateParams) (models.User, error)
	PartialUpdate(ctx context.Context, params *PartialUpdateParams) (models.User, error)
	UpdateAvatar(ctx context.Context, id int, avatar string) error

	Delete(ctx context.Context, id int) error

	Exists(ctx context.Context, id int) (bool, error)
}
