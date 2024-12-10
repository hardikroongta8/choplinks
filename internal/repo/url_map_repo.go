package repo

import (
	"github.com/hardikroongta8/choplinks/internal/model"
	"gorm.io/gorm"
)

type UrlMapsRepo struct {
	DB *gorm.DB
}

func (repo *UrlMapsRepo) InsertUrlMap(urlMap *model.UrlMap) error {
	res := repo.DB.Create(urlMap)
	if res.Error == nil {
		return res.Error
	}
	return nil
}

func (repo *UrlMapsRepo) FindUrlMap(shortenedURLPath string) (*model.UrlMap, error) {
	var urlMap model.UrlMap
	res := repo.DB.First(&urlMap, "shortened_url_path = ?", shortenedURLPath)
	if res.Error != nil {
		return nil, res.Error
	}
	return &urlMap, nil
}
