package usecase

import (
	"context"
	"github.com/SlavaShagalov/prospeech-backend/internal/files"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/config"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/constants"
	pkgErrors "github.com/SlavaShagalov/prospeech-backend/internal/pkg/errors"
	"github.com/SlavaShagalov/prospeech-backend/internal/users"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	avatarsFolder = "avatars"
)

type usecase struct {
	usersRepo users.Repository
	imgRepo   files.Repository
}

func New(rep users.Repository, imgRepo files.Repository) users.Usecase {
	return &usecase{
		usersRepo: rep,
		imgRepo:   imgRepo,
	}
}

func (uc *usecase) List() ([]models.User, error) {
	return uc.usersRepo.List(context.TODO())
}

func (uc *usecase) Get(id int64) (models.User, error) {
	return uc.usersRepo.Get(context.TODO(), id)
}

func (uc *usecase) GetByUsername(username string) (models.User, error) {
	return uc.usersRepo.GetByUsername(context.TODO(), username)
}

func (uc *usecase) FullUpdate(params *users.FullUpdateParams) (models.User, error) {
	if err := validateUsername(params.Username); err != nil {
		return models.User{}, err
	} else if err = validateName(params.Name); err != nil {
		return models.User{}, err
	}

	_, err := uc.usersRepo.GetByUsername(context.TODO(), params.Username)
	if !errors.Is(err, pkgErrors.ErrUserNotFound) {
		if err != nil {
			return models.User{}, err
		}
		return models.User{}, pkgErrors.ErrUserAlreadyExists
	}

	return uc.usersRepo.FullUpdate(context.TODO(), params)
}

func (uc *usecase) PartialUpdate(params *users.PartialUpdateParams) (models.User, error) {
	if params.UpdateUsername {
		if err := validateUsername(params.Username); err != nil {
			return models.User{}, err
		}

		user, err := uc.usersRepo.GetByUsername(context.TODO(), params.Username)
		if !errors.Is(err, pkgErrors.ErrUserNotFound) && user.ID != params.ID {
			if err != nil {
				return models.User{}, err
			}
			return models.User{}, pkgErrors.ErrUserAlreadyExists
		}
	} else if params.UpdateName {
		if err := validateName(params.Name); err != nil {
			return models.User{}, err
		}
	}

	return uc.usersRepo.PartialUpdate(context.TODO(), params)
}

func (uc *usecase) UpdateAvatar(id int64, imgData []byte, filename string) (*models.User, error) {
	//user, err := uc.usersRepo.Get(context.TODO(), id)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if user.Avatar == nil {
	//	imgName := avatarsFolder + "/" + uuid.NewString() + filepath.Ext(filename)
	//	imgPath, err := uc.imgRepo.Create(imgName, imgData)
	//	if err == nil {
	//		err = uc.usersRepo.UpdateAvatar(context.TODO(), id, imgPath)
	//		if err == nil {
	//			user.Avatar = &imgPath
	//		}
	//	}
	//} else {
	//	err = uc.imgRepo.Update(*user.Avatar, imgData)
	//}
	//
	//return &user, err
	return nil, nil
}

func (uc *usecase) Delete(id int64) error {
	return uc.usersRepo.Delete(context.TODO(), id)
}

func validateUsername(username string) error {
	if len(username) < viper.GetInt(config.MinUsernameLen) {
		return pkgErrors.ErrTooShortUsername
	} else if len(username) > viper.GetInt(config.MaxUsernameLen) {
		return pkgErrors.ErrTooLongUsername
	}
	return nil
}

func validateName(name string) error {
	if len(name) < constants.MinNameLen {
		return pkgErrors.ErrEmptyName
	} else if len(name) > constants.MaxNameLen {
		return pkgErrors.ErrTooLongName
	}
	return nil
}
