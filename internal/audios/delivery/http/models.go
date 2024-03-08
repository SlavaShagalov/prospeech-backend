package http

import (
	"github.com/SlavaShagalov/prospeech-backend/internal/models"
	"time"
)

// API requests
type partialUpdateRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// API responses
type listResponse struct {
	Audios []models.Audio `json:"audios"`
}

func newListResponse(audios []models.Audio) *listResponse {
	return &listResponse{
		Audios: audios,
	}
}

type createResponse struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newCreateResponse(audio *models.Audio) *createResponse {
	return &createResponse{
		ID:        audio.ID,
		UserID:    audio.UserID,
		Title:     audio.Title,
		URL:       audio.URL,
		CreatedAt: audio.CreatedAt,
		UpdatedAt: audio.UpdatedAt,
	}
}

//type getResponse struct {
//	ID          int       `json:"id"`
//	Title       string    `json:"title"`
//	Description string    `json:"description"`
//	Background  *string   `json:"background"`
//	CreatedAt   time.Time `json:"created_at"`
//	UpdatedAt   time.Time `json:"updated_at"`
//}
//
//func newGetResponse(audio *models.Audio) *getResponse {
//	return &getResponse{
//		ID:          audio.ID,
//		Title:       audio.Title,
//		Description: audio.Description,
//		Background:  audio.Background,
//		CreatedAt:   audio.CreatedAt,
//		UpdatedAt:   audio.UpdatedAt,
//	}
//}
