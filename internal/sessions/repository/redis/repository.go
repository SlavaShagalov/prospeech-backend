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

func (repo *repository) Create(ctx context.Context, userID int) (string, error) {
	authToken := strconv.Itoa(userID) + "$" + uuid.New().String()

	err := repo.rdb.HSet(ctx, strconv.Itoa(userID), authToken, []byte{}).Err()
	if err != nil {
		repo.log.Error("Failed to set key-value in Redis", zap.Error(err), zap.Int("user_id", userID),
			zap.String("auth_token", authToken))
		repo.rdb.Expire(ctx, strconv.Itoa(userID), 5*time.Hour)
		return "", err
	}

	return authToken, nil
}

func (repo *repository) Get(ctx context.Context, userID int, authToken string) (int, error) {
	err := repo.rdb.HGet(ctx, strconv.Itoa(userID), authToken).Err()
	if err != nil {
		repo.log.Info("Failed to get session", zap.Error(err), zap.Int("user_id", userID),
			zap.String("token", authToken))
		return 0, pkgErrors.ErrSessionNotFound
	}

	return userID, nil
}

func (repo *repository) Delete(ctx context.Context, userID int, authToken string) error {
	if err := repo.rdb.HGet(ctx, strconv.Itoa(userID), authToken).Err(); err != nil {
		repo.log.Info("Failed to delete session", zap.Error(err), zap.Int("user_id", userID),
			zap.String("token", authToken))
		return pkgErrors.ErrSessionNotFound
	}

	repo.rdb.HDel(ctx, strconv.Itoa(userID), authToken)
	return nil
}
