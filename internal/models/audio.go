package models

import "time"

type Audio struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	Text       string `json:"text"`
	Words      []string
	StartTimes []float32
	EndTimes   []float32
	Duration   time.Duration
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
