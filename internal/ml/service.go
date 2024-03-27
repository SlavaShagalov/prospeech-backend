package ml

import (
	"bytes"
	"encoding/json"
	"github.com/SlavaShagalov/prospeech-backend/internal/pkg/constants"
	"go.uber.org/zap"
	"io"
	"net/http"
)

const (
	wav2VecUrl     = "http://ml:8080/predictions/Wav2Vec2"
	improveTextUrl = "http://ml:8080/predictions/Seq2Seq"
)

type Service struct {
	log *zap.Logger
}

func New(log *zap.Logger) *Service {
	return &Service{
		log: log,
	}
}

type Data struct {
	Words       []string  `json:"words"`
	StartTimes  []float64 `json:"word_start_times"`
	EndTimes    []float64 `json:"word_end_times"`
	WordsPerMin uint      `json:"words_per_minute"`
}

func (s *Service) Wav2Vec(wavData []byte) (*Data, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, wav2VecUrl, bytes.NewBuffer(wavData))
	if err != nil {
		s.log.Error("Wav2Vec: failed to create HTTP request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Content-Type", "audio/basic")
	resp, err := client.Do(req)
	if err != nil {
		s.log.Error("Wav2Vec: request failed", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.Error(constants.FailedReadRequestBody, zap.Error(err))
		return nil, err
	}

	result := new(Data)
	err = json.Unmarshal(body, result)
	if err != nil {
		s.log.Error(constants.FailedMarshalBody, zap.Error(err))
		return nil, err
	}

	return result, nil
}

func (s *Service) ImproveText(badText string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, improveTextUrl, bytes.NewBuffer([]byte(badText)))
	if err != nil {
		s.log.Error("ImproveText: failed to create HTTP request", zap.Error(err))
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		s.log.Error("ImproveText: request failed", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.Error(constants.FailedReadRequestBody, zap.Error(err))
		return "", err
	}

	var result map[string][]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		s.log.Error(constants.FailedMarshalBody, zap.Error(err))
		return "", err
	}

	return result["transcription"][0], nil
}
