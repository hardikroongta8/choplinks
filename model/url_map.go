package model

import "time"

type URLMap struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateURLMapReqBody struct {
	OriginalUrl string `json:"original_url"`
}
