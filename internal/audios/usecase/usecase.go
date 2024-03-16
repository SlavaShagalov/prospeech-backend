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
	"time"
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

func runML(filename string) {
	cmd := exec.Command("python3", "/bin/ml/main.py", filename)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return
	}
}

type Data struct {
	Words      []string      `json:"words"`
	StartTimes []float64     `json:"start_times"`
	EndTimes   []float64     `json:"end_times"`
	Duration   time.Duration `json:"duration"`
}

func analyze(file *pFiles.File) (string, error) {
	log.Println("Start processing " + file.Name)
	log.Println("End processing " + file.Name)
	return "Hello from ML!", nil
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

	text, err := analyze(&params.File)
	if err != nil {
		return nil, err
	}

	//var data Data
	//err = os.WriteFile("/data/speech."+filepath.Ext(params.File.Name), params.File.Data, 0777)
	//if err != nil {
	//	fmt.Println("Ошибка при записи в файл:", err)
	//} else {
	//	runML("/data/speech." + filepath.Ext(params.File.Name))
	//
	//	file, err := os.Open("/data/speech.json")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	defer file.Close()
	//
	//	err = json.NewDecoder(file).Decode(&data)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	repoParams := audiosRepo.CreateParams{
		UserID: params.UserID,
		Title:  "Untitled Speech",
		URL:    url,
		Text:   text,
	}
	audio, err := uc.repo.Create(ctx, &repoParams)
	log.Println(audio)
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
