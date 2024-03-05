package pgx

import (
	"context"
	"database/sql"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/constants"
	pkgErrors "github.com/SlavaShagalov/prospeech-backend/internal/pkg/errors"
	pkgUsers "github.com/SlavaShagalov/prospeech-backend/internal/users"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type repository struct {
	pool *pgxpool.Pool
	log  *zap.Logger
}

func New(pool *pgxpool.Pool, log *zap.Logger) pkgUsers.Repository {
	return &repository{
		pool: pool,
		log:  log,
	}
}

const createCmd = `
	INSERT INTO users (name, username, email, hashed_password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, username, hashed_password, email, name, avatar, created_at, updated_at;`

func (repo *repository) Create(ctx context.Context, params *pkgUsers.CreateParams) (models.User, error) {
	row := repo.pool.QueryRow(ctx, createCmd, params.Name, params.Username, params.Email, params.HashedPassword)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", createCmd),
			zap.Any("create_params", params))
		return models.User{}, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}

	repo.log.Debug("User created", zap.Int("id", user.ID), zap.String("username", user.Username))
	return user, nil
}

const listCmd = `
	SELECT id, username, hashed_password, email, name, avatar, created_at, updated_at
	FROM users;`

func (repo *repository) List(ctx context.Context) ([]models.User, error) {
	rows, err := repo.pool.Query(ctx, listCmd)
	if err != nil {
		repo.log.Error(constants.DBError, zap.Error(err), zap.String("sql_query", listCmd))
		return nil, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}
	defer rows.Close()

	users := []models.User{}
	var user models.User
	var avatar sql.NullString
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Name,
			&avatar,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", listCmd))
			return nil, errors.Wrap(pkgErrors.ErrDb, err.Error())
		}

		if avatar.Valid {
			user.Avatar = &avatar.String
		} else {
			user.Avatar = nil
		}

		users = append(users, user)
	}

	return users, nil
}

const getCmd = `
	SELECT id, username, hashed_password, email, name, avatar, created_at, updated_at
	FROM users
	WHERE id = $1;`

func (repo *repository) Get(ctx context.Context, id int) (models.User, error) {
	row := repo.pool.QueryRow(ctx, getCmd, id)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, errors.Wrap(pkgErrors.ErrUserNotFound, err.Error())
		}

		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", getCmd),
			zap.Int("id", id))
		return models.User{}, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}

	return user, nil
}

const getByUsernameCmd = `
	SELECT id, username, hashed_password, email, name, avatar, created_at, updated_at
	FROM users
	WHERE username = $1;`

func (repo *repository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	row := repo.pool.QueryRow(ctx, getByUsernameCmd, username)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, pkgErrors.ErrUserNotFound
		}

		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", getByUsernameCmd),
			zap.String("username", username))
		return models.User{}, pkgErrors.ErrDb
	}

	return user, nil
}

const fullUpdateCmd = `
	UPDATE users
	SET username = $1,
	    email    = $2,
		name     = $3
	WHERE id = $4
	RETURNING id, username, hashed_password, email, name, avatar, created_at, updated_at;`

func (repo *repository) FullUpdate(ctx context.Context, params *pkgUsers.FullUpdateParams) (models.User, error) {
	row := repo.pool.QueryRow(ctx, fullUpdateCmd, params.Username, params.Email, params.Name, params.ID)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", fullUpdateCmd),
			zap.Any("params", params))
		return models.User{}, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}

	repo.log.Debug("User full updated", zap.Any("user", user))
	return user, nil
}

const partialUpdateCmd = `
	UPDATE users
	SET username = CASE WHEN $1::boolean THEN $2 ELSE username END,
		email    = CASE WHEN $3::boolean THEN $4 ELSE email END,
		name     = CASE WHEN $5::boolean THEN $6 ELSE name END
	WHERE id = $7
	RETURNING id, username, hashed_password, email, name, avatar, created_at, updated_at;`

func (repo *repository) PartialUpdate(ctx context.Context, params *pkgUsers.PartialUpdateParams) (models.User, error) {
	row := repo.pool.QueryRow(ctx, partialUpdateCmd,
		params.UpdateUsername,
		params.Username,
		params.UpdateEmail,
		params.Email,
		params.UpdateName,
		params.Name,
		params.ID,
	)

	var user models.User
	err := scanUser(row, &user)
	if err != nil {
		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", partialUpdateCmd),
			zap.Any("params", params))
		return models.User{}, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}

	repo.log.Debug("User partial updated", zap.Any("user", user))
	return user, nil
}

const updateAvatarCmd = `
	UPDATE users
	SET avatar = $1
	WHERE id = $2;`

func (repo *repository) UpdateAvatar(ctx context.Context, id int, avatar string) error {
	result, err := repo.pool.Exec(ctx, updateAvatarCmd, avatar, id)
	if err != nil {
		repo.log.Error(constants.DBError, zap.Error(err), zap.String("sql", updateAvatarCmd),
			zap.Int("id", id))
		return pkgErrors.ErrDb
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return pkgErrors.ErrUserNotFound
	}

	repo.log.Debug("Avatar updated", zap.Int("id", id))
	return nil
}

const deleteCmd = `
	DELETE FROM users 
	WHERE id = $1;`

func (repo *repository) Delete(ctx context.Context, id int) error {
	result, err := repo.pool.Exec(ctx, deleteCmd, id)
	if err != nil {
		repo.log.Error(constants.DBError, zap.Error(err), zap.String("sql_query", deleteCmd),
			zap.Int("id", id))
		return errors.Wrap(pkgErrors.ErrDb, err.Error())
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return pkgErrors.ErrUserNotFound
	}

	repo.log.Debug("User deleted", zap.Int("id", id))
	return nil
}

const existsCmd = `
	SELECT EXISTS(SELECT id
					FROM users
					WHERE id = $1) AS exists;`

func (repo *repository) Exists(ctx context.Context, userID int) (bool, error) {
	row := repo.pool.QueryRow(ctx, existsCmd, userID)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		repo.log.Error(constants.DBScanError, zap.Error(err), zap.String("sql_query", existsCmd),
			zap.Int("user_id", userID))
		return false, errors.Wrap(pkgErrors.ErrDb, err.Error())
	}
	return exists, nil
}

func scanUser(row pgx.Row, user *models.User) error {
	avatar := new(sql.NullString)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		avatar,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	if avatar.Valid {
		user.Avatar = &avatar.String
	} else {
		user.Avatar = nil
	}

	return nil
}
