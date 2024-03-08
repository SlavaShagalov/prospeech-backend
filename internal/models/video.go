package models

import "time"

type Video struct {
	ID        int64
	UserID    int64
	Title     string
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
