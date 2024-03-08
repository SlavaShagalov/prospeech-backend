package std

import (
	"context"
	"database/sql"
	pAudiosRepo "github.com/SlavaShagalov/prospeech-backend/internal/audios/repository"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/constants"
	pErrors "github.com/SlavaShagalov/prospeech-backend/internal/pkg/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type repository struct {
	pool *pgxpool.Pool
	log  *zap.Logger
}

func New(pool *pgxpool.Pool, log *zap.Logger) pAudiosRepo.Repository {
	return &repository{
		pool: pool,
		log:  log,
	}
}

const createCmd = `
	INSERT INTO audios (user_id, title, URL) 
	VALUES ($1, $2, $3)
	RETURNING id, user_id, title, URL, created_at, updated_at;`

func (repo *repository) Create(ctx context.Context, params *pAudiosRepo.CreateParams) (*models.Audio, error) {
	row := repo.pool.QueryRow(ctx, createCmd, params.UserID, params.Title, params.URL)

	audio := new(models.Audio)
	err := scanAudio(row, audio)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if !ok {
			repo.log.Error("Cannot convert err to pq.Error", zap.Error(err))
			return nil, errors.Wrap(pErrors.ErrDb, err.Error())
		}
		if pgErr.Constraint == "audios_user_id_fkey" {
			return nil, errors.Wrap(pErrors.ErrAudioNotFound, err.Error())
		}

		repo.log.Error(constants.DBScanError, zap.Error(err), zap.Any("params", params))
		return nil, errors.Wrap(pErrors.ErrDb, err.Error())
	}

	repo.log.Debug("New audio created", zap.Int64("id", audio.ID))
	return audio, nil
}

const listCmd = `
	SELECT id, user_id, title, url, created_at, updated_at
	FROM audios
	WHERE user_id = $1;`

func (repo *repository) List(ctx context.Context, userID int64) ([]models.Audio, error) {
	rows, err := repo.pool.Query(ctx, listCmd, userID)
	if err != nil {
		repo.log.Error(constants.DBError, zap.Error(err), zap.Int64("user_id", userID))
		return nil, errors.Wrap(pErrors.ErrDb, err.Error())
	}
	defer rows.Close()

	audios := []models.Audio{}
	var audio models.Audio
	for rows.Next() {
		err = rows.Scan(
			&audio.ID,
			&audio.UserID,
			&audio.Title,
			&audio.URL,
			&audio.CreatedAt,
			&audio.UpdatedAt,
		)
		if err != nil {
			repo.log.Error(constants.DBScanError, zap.Error(err), zap.Int64("user_id", userID))
			return nil, errors.Wrap(pErrors.ErrDb, err.Error())
		}

		audios = append(audios, audio)
	}

	return audios, nil
}

const getCmd = `
	SELECT id, user_id, title, url, created_at, updated_at
	FROM audios
	WHERE id = $1;`

func (repo *repository) Get(ctx context.Context, id int64) (*models.Audio, error) {
	row := repo.pool.QueryRow(ctx, getCmd, id)

	var audio *models.Audio
	err := scanAudio(row, audio)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(pErrors.ErrAudioNotFound, err.Error())
		}

		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", getCmd),
			zap.Int64("id", id))
		return nil, errors.Wrap(pErrors.ErrDb, err.Error())
	}

	return audio, nil
}

//const partialUpdateCmd = `
//	UPDATE audios
//	SET title        = CASE WHEN $1::boolean THEN $2 ELSE title END,
//		description  = CASE WHEN $3::boolean THEN $4 ELSE description END,
//		user_id = CASE WHEN $5::boolean THEN $6 ELSE user_id END
//	WHERE id = $7
//	RETURNING id, user_id, title, description, background, created_at, updated_at;`
//
//func (repo *repository) PartialUpdate(ctx context.Context, params *pkgAudios.PartialUpdateParams) (models.Audio, error) {
//	row := repo.pool.QueryRow(ctx, partialUpdateCmd,
//		params.UpdateTitle,
//		params.Title,
//		params.UpdateDescription,
//		params.Description,
//		params.UpdateUserID,
//		params.UserID,
//		params.ID,
//	)
//
//	var audio models.Audio
//	err := scanAudio(row, &audio)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, errors.Wrap(pErrors.ErrAudioNotFound, err.Error())
//		}
//
//		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", partialUpdateCmd),
//			zap.Any("params", params))
//		return nil, errors.Wrap(pErrors.ErrDb, err.Error())
//	}
//
//	repo.log.Debug("Audio partial updated", zap.Any("audio", audio))
//	return audio, nil
//}

const deleteCmd = `
	DELETE FROM audios 
	WHERE id = $1;`

func (repo *repository) Delete(ctx context.Context, id int64) error {
	result, err := repo.pool.Exec(ctx, deleteCmd, id)
	if err != nil {
		repo.log.Error(constants.DBError, zap.Error(err), zap.String("sql_query", deleteCmd),
			zap.Int64("id", id))
		return errors.Wrap(pErrors.ErrDb, err.Error())
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return pErrors.ErrAudioNotFound
	}

	repo.log.Debug("Audio deleted", zap.Int64("id", id))
	return nil
}

func scanAudio(row pgx.Row, audio *models.Audio) error {
	err := row.Scan(
		&audio.ID,
		&audio.UserID,
		&audio.Title,
		&audio.URL,
		&audio.CreatedAt,
		&audio.UpdatedAt,
	)
	return err
}
