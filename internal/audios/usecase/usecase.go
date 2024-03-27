package usecase

import (
	"context"
	pAudios "github.com/SlavaShagalov/prospeech-backend/internal/audios"
	"github.com/SlavaShagalov/prospeech-backend/internal/audios/repository"
	audiosRepo "github.com/SlavaShagalov/prospeech-backend/internal/audios/repository"
	"github.com/SlavaShagalov/prospeech-backend/internal/files"
	pFiles "github.com/SlavaShagalov/prospeech-backend/internal/files"
	"github.com/SlavaShagalov/prospeech-backend/internal/ml"
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/convert"
	"github.com/SlavaShagalov/prospeech-backend/internal/users"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	audiosFolder = "audios"
)

type usecase struct {
	repo      repository.Repository
	filesRepo files.Repository
	usersRepo users.Repository
	mlServ    *ml.Service
	log       *zap.Logger
}

func New(repo repository.Repository, filesRepo files.Repository, usersRepo users.Repository, mlServ *ml.Service, log *zap.Logger) pAudios.Usecase {
	return &usecase{
		repo:      repo,
		filesRepo: filesRepo,
		usersRepo: usersRepo,
		mlServ:    mlServ,
		log:       log,
	}
}

func (uc *usecase) Create(ctx context.Context, params *pAudios.CreateParams) (*models.Audio, error) {
	fileS3 := pFiles.File{
		Name: audiosFolder + "/" + uuid.NewString() + filepath.Ext(params.File.Name),
		Data: params.File.Data,
	}
	url, err := uc.filesRepo.Create(ctx, &fileS3)
	if err != nil {
		return nil, err
	}

	wavData, err := convert.MP4ToWAV(params.File.Data)
	if err != nil {
		uc.log.Error("Failed to convert MP4 to WAV", zap.Error(err))
		return nil, err
	}

	mlData, err := uc.mlServ.Wav2Vec(wavData)
	if err != nil {
		return nil, err
	}

	badText := strings.Join(mlData.Words, " ")
	improvedText, err := uc.mlServ.ImproveText(badText)
	if err != nil {
		return nil, err
	}

	title := "Выступление"
	curCount, err := uc.usersRepo.UpdateUntitledSpeechesCount(ctx, params.UserID)
	if err == nil {
		title += " " + strconv.Itoa(curCount)
	}

	repoParams := audiosRepo.CreateParams{
		UserID:      params.UserID,
		Title:       title,
		URL:         url,
		Text:        improvedText,
		Words:       mlData.Words,
		StartTimes:  mlData.StartTimes,
		EndTimes:    mlData.EndTimes,
		WordsPerMin: mlData.WordsPerMin,
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

func (uc *usecase) PartialUpdate(ctx context.Context, params *audiosRepo.PartialUpdateParams) (*models.Audio, error) {
	return uc.repo.PartialUpdate(ctx, params)
}

func (uc *usecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
