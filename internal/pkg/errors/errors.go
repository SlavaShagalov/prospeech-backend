package errors

import (
	"errors"
	"fmt"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/constants"
)

var (
	// Common repository
	ErrDb = errors.New("database error")

	// Users
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrTooShortUsername  = errors.New(fmt.Sprintf("username must be at least %d characters",
		constants.MinUsernameLen))
	ErrTooLongUsername = errors.New(fmt.Sprintf("username must be no more than %d characters",
		constants.MaxUsernameLen))
	ErrEmptyName   = errors.New("name must not be empty")
	ErrTooLongName = errors.New(fmt.Sprintf("name must be no more than %d characters", constants.MaxNameLen))

	// Workspaces
	ErrAudioNotFound = errors.New("audio not found")

	// Auth
	ErrWrongLoginOrPassword = errors.New("wrong login or password")
	ErrGetHashedPassword    = errors.New("get hashed password error")
	// ErrSessionStorage       = errors.New("session storage error")
	ErrSessionNotFound = errors.New("session not found")

	// HTTP
	ErrReadBody         = errors.New("read request body error")
	ErrBadSessionCookie = errors.New("bad session cookie")
)
