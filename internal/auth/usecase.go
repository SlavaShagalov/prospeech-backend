package auth

import (
	"context"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
)

type SignInParams struct {
	Username string
	Password string
}

type SignUpParams struct {
	Name     string
	Username string
	Email    string
	Password string
}

type Usecase interface {
	SignIn(ctx context.Context, params *SignInParams) (models.User, string, error)
	SignUp(ctx context.Context, params *SignUpParams) (models.User, string, error)
	CheckAuth(ctx context.Context, userID int64, authToken string) (int64, error)
	Logout(ctx context.Context, userID int64, authToken string) error
}
