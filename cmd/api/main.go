package main

import (
	"context"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/config"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/storages/postgres"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	mw "github.com/SlavaShagalov/prospeech-backend/internal/middleware"
	pHasher "github.com/SlavaShagalov/prospeech-backend/internal/pkg/hasher/bcrypt"
	pLog "github.com/SlavaShagalov/prospeech-backend/internal/pkg/log/zap"
	pStorages "github.com/SlavaShagalov/prospeech-backend/internal/pkg/storages"

	audiosRepository "github.com/SlavaShagalov/prospeech-backend/internal/audios/repository/pgx"
	filesRepository "github.com/SlavaShagalov/prospeech-backend/internal/files/repository/s3"
	sessionsRepository "github.com/SlavaShagalov/prospeech-backend/internal/sessions/repository/redis"
	usersRepository "github.com/SlavaShagalov/prospeech-backend/internal/users/repository/pgx"

	audiosUsecase "github.com/SlavaShagalov/prospeech-backend/internal/audios/usecase"
	authUsecase "github.com/SlavaShagalov/prospeech-backend/internal/auth/usecase"
	usersUsecase "github.com/SlavaShagalov/prospeech-backend/internal/users/usecase"

	audiosDel "github.com/SlavaShagalov/prospeech-backend/internal/audios/delivery/http"
	authDel "github.com/SlavaShagalov/prospeech-backend/internal/auth/delivery/http"
	usersDel "github.com/SlavaShagalov/prospeech-backend/internal/users/delivery/http"
)

func main() {
	ctx := context.Background()

	// ===== Configuration =====
	config.SetDefaultPostgresConfig()
	config.SetDefaultRedisConfig()
	config.SetDefaultS3Config()
	config.SetDefaultValidationConfig()
	viper.SetConfigName("api")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Failed to read configuration: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Configuration read successfully")

	// ===== Logger =====
	logger, logfile, err := pLog.NewProdLogger("/logs/" + viper.GetString(config.ServerName) + ".log")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Println(err)
		}
		err = logfile.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	logger.Info("API server starting...")

	// ===== Data Storage =====
	pgxPool, err := postgres.NewPgx(logger)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		pgxPool.Close()
		logger.Info("Postgres connection closed")
	}()

	// ===== Sessions Storage =====
	redisClient, err := pStorages.NewRedis(logger, ctx)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		err = redisClient.Close()
		if err != nil {
			logger.Error("Failed to close Redis client", zap.Error(err))
		}
		logger.Info("Redis client closed")
	}()

	// ===== S3 =====
	s3Client, err := pStorages.NewS3(logger)
	if err != nil {
		os.Exit(1)
	}

	// ===== Hasher =====
	hasher := pHasher.New()

	// ===== Repositories =====
	usersRepo := usersRepository.New(pgxPool, logger)
	filesRepo := filesRepository.New(s3Client, logger)
	audiosRepo := audiosRepository.New(pgxPool, logger)
	sessionsRepo := sessionsRepository.New(redisClient, logger)

	// ===== Usecases =====
	authUC := authUsecase.New(usersRepo, sessionsRepo, hasher, logger)
	usersUC := usersUsecase.New(usersRepo, filesRepo)
	audiosUC := audiosUsecase.New(audiosRepo, filesRepo)

	// ===== Middleware =====
	checkAuth := mw.NewCheckAuth(authUC, logger)
	accessLog := mw.NewAccessLog(logger)
	cors := mw.NewCors()

	router := mux.NewRouter()

	// ===== Delivery =====
	authDel.RegisterHandlers(router, authUC, usersUC, logger, checkAuth)
	usersDel.RegisterHandlers(router, usersUC, logger, checkAuth)
	audiosDel.RegisterHandlers(router, audiosUC, logger, checkAuth)

	// ===== Router =====
	server := http.Server{
		Addr:    ":" + viper.GetString(config.ServerPort),
		Handler: accessLog(cors(router)),
	}

	// ===== Start =====
	logger.Info("API server started", zap.String("port", viper.GetString(config.ServerPort)))
	if err = server.ListenAndServe(); err != nil {
		logger.Error("API server stopped", zap.Error(err))
	}
}
