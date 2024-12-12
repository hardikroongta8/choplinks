package model

import (
	"time"
)

type UrlMap struct {
	ShortenedURLPath string `gorm:"primaryKey"`
	OriginalURL      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
