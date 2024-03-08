package sessions

import "context"

type Repository interface {
	Create(ctx context.Context, userID int64) (string, error)
	Get(ctx context.Context, userID int64, authToken string) (int64, error)
	Delete(ctx context.Context, userID int64, authToken string) error
}
