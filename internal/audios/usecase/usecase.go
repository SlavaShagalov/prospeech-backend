package usecase

import (
	"context"
	pAudios "github.com/SlavaShagalov/prospeech-backend/internal/audios"
	"github.com/SlavaShagalov/prospeech-backend/internal/audios/repository"
	audiosRepo "github.com/SlavaShagalov/prospeech-backend/internal/audios/repository"
	"github.com/SlavaShagalov/prospeech-backend/internal/files"
	pFiles "github.com/SlavaShagalov/prospeech-backend/internal/files"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
	"github.com/google/uuid"
	"log"
	"os/exec"
	"path/filepath"
)

const (
	audiosFolder = "audios"
)

type usecase struct {
	repo      repository.Repository
	filesRepo files.Repository
}

func New(repo repository.Repository, filesRepo files.Repository) pAudios.Usecase {
	return &usecase{
		repo:      repo,
		filesRepo: filesRepo,
	}
}

func runML() {
	cmd := exec.Command("python3", "/bin/ml/main.py")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func (uc *usecase) Create(ctx context.Context, params *pAudios.CreateParams) (*models.Audio, error) {
	//runML()

	file := pFiles.File{
		Name: audiosFolder + "/" + uuid.NewString() + filepath.Ext(params.File.Name),
		Data: params.File.Data,
	}
	url, err := uc.filesRepo.Create(ctx, &file)
	if err != nil {
		return nil, err
	}
	//url := file.Name

	repoParams := audiosRepo.CreateParams{
		UserID: params.UserID,
		Title:  params.File.Name,
		URL:    url,
	}
	audio, err := uc.repo.Create(ctx, &repoParams)
	return audio, err
}

func (uc *usecase) List(ctx context.Context, userID int64) ([]models.Audio, error) {
	return uc.repo.List(ctx, userID)
}

func (uc *usecase) Get(ctx context.Context, id int64) (*models.Audio, error) {
	return uc.repo.Get(ctx, id)
}

//func (uc *usecase) PartialUpdate(ctx context.Context, params *pAudios.PartialUpdateParams) (models.audio, error) {
//	ctx, span := opentel.Tracer.Start(ctx, componentName+" "+"PartialUpdate")
//	defer span.End()
//
//	return uc.repo.PartialUpdate(ctx, params)
//}
//
//func (uc *usecase) UpdateBackground(ctx context.Context, id int, imgData []byte, filename string) (*models.audio, error) {
//	ctx, span := opentel.Tracer.Start(ctx, componentName+" "+"UpdateBackground")
//	defer span.End()
//
//	audio, err := uc.repo.Get(ctx, id)
//	if err != nil {
//		return nil, err
//	}
//
//	if audio.Background == nil {
//		imgName := backgroundsFolder + "/" + uuid.NewString() + filepath.Ext(filename)
//		imgPath, err := uc.filesRepo.Create(imgName, imgData)
//		if err == nil {
//			err = uc.repo.UpdateBackground(ctx, id, imgPath)
//			if err == nil {
//				audio.Background = &imgPath
//			}
//		}
//	} else {
//		err = uc.filesRepo.Update(*audio.Background, imgData)
//	}
//
//	return &audio, err
//}
//
//func (uc *usecase) Delete(ctx context.Context, id int) error {
//	ctx, span := opentel.Tracer.Start(ctx, componentName+" "+"Delete")
//	defer span.End()
//
//	return uc.repo.Delete(ctx, id)
//}
