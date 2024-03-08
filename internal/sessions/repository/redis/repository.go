package redis

import (
	"context"
	"go.uber.org/zap"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	pkgErrors "github.com/SlavaShagalov/prospeech-backend/internal/pkg/errors"
	pkgSessions "github.com/SlavaShagalov/prospeech-backend/internal/sessions"
)

type repository struct {
	rdb *redis.Client
	log *zap.Logger
}

func New(rdb *redis.Client, log *zap.Logger) pkgSessions.Repository {
	return &repository{
		rdb: rdb,
		log: log,
	}
}

func (repo *repository) Create(ctx context.Context, userID int64) (string, error) {
	authToken := strconv.FormatInt(userID, 10) + "$" + uuid.New().String()

	err := repo.rdb.HSet(ctx, strconv.FormatInt(userID, 10), authToken, []byte{}).Err()
	if err != nil {
		repo.log.Error("Failed to set key-value in Redis", zap.Error(err), zap.Int64("user_id", userID),
			zap.String("auth_token", authToken))
		repo.rdb.Expire(ctx, strconv.FormatInt(userID, 10), 5*time.Hour)
		return "", err
	}

	return authToken, nil
}

func (repo *repository) Get(ctx context.Context, userID int64, authToken string) (int64, error) {
	err := repo.rdb.HGet(ctx, strconv.FormatInt(userID, 10), authToken).Err()
	if err != nil {
		repo.log.Info("Failed to get session", zap.Error(err), zap.Int64("user_id", userID),
			zap.String("token", authToken))
		return 0, pkgErrors.ErrSessionNotFound
	}

	return userID, nil
}

func (repo *repository) Delete(ctx context.Context, userID int64, authToken string) error {
	if err := repo.rdb.HGet(ctx, strconv.FormatInt(userID, 10), authToken).Err(); err != nil {
		repo.log.Info("Failed to delete session", zap.Error(err), zap.Int64("user_id", userID),
			zap.String("token", authToken))
		return pkgErrors.ErrSessionNotFound
	}

	repo.rdb.HDel(ctx, strconv.FormatInt(userID, 10), authToken)
	return nil
}
