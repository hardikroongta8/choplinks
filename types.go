package main

import "time"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type URLMap struct {
	ShortenedURLPath string    `json:"shortened_url_path"`
	OriginalURL      string    `json:"original_url"`
	UserID           string    `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreateUserReqBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateURLMapReqBody struct {
	OriginalUrl string `json:"original_url"`
	UserID      string `json:"user_id"`
}

type GetURLMapsReqBody struct {
	UserID string `json:"user_id"`
}
