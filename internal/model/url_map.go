package model

import (
	"gorm.io/gorm"
)

type UrlMap struct {
	ShortenedURLPath string `gorm:"primaryKey"`
	OriginalURL      string
	gorm.Model
}
