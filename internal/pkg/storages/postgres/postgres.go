package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/config"
)

func NewStd(log *zap.Logger) (*sql.DB, error) {
	log.Info("Connecting to Postgres...",
		zap.String("host", viper.GetString(config.PostgresHost)),
		zap.Int("port", viper.GetInt(config.PostgresPort)),
		zap.String("dbname", viper.GetString(config.PostgresDB)),
	)

	params := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString(config.PostgresHost),
		viper.GetInt(config.PostgresPort),
		viper.GetString(config.PostgresUser),
		viper.GetString(config.PostgresDB),
		viper.GetString(config.PostgresPassword),
		viper.GetString(config.PostgresSSLMode),
	)

	db, err := sql.Open("postgres", params)
	if err != nil {
		log.Error("Failed to create Postgres connection", zap.Error(err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error("Failed to connect to Postgres", zap.Error(err))
		return nil, err
	}

	log.Info("Postgres connection created successfully")
	return db, nil
}

func NewPgx(log *zap.Logger) (*pgxpool.Pool, error) {
	log.Info("Connecting to Postgres PGX...",
		zap.String("host", viper.GetString(config.PostgresHost)),
		zap.Int("port", viper.GetInt(config.PostgresPort)),
		zap.String("dbname", viper.GetString(config.PostgresDB)),
	)

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=100",
		viper.GetString(config.PostgresUser),
		viper.GetString(config.PostgresPassword),
		viper.GetString(config.PostgresHost),
		strconv.Itoa(viper.GetInt(config.PostgresPort)),
		viper.GetString(config.PostgresDB),
	)

	conf, _ := pgxpool.ParseConfig(connString)
	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		log.Error("Failed to connect to db PGX", zap.Error(err))
		return nil, err
	}

	log.Info("Postgres PGX connection created successfully")
	return pool, nil
}
