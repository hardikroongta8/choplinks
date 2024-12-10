package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hardikroongta8/choplinks/internal/model"
	"github.com/hardikroongta8/choplinks/internal/repo"
	"github.com/hardikroongta8/choplinks/pkg/config"
	"github.com/hardikroongta8/choplinks/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type URLMapService struct {
	DB *gorm.DB
}

func (service *URLMapService) CreateUrlMap(c *gin.Context) {
	c.Request.Context()
	type Url struct {
		Url string `json:"url"`
	}
	var ogUrl Url
	err := c.ShouldBindBodyWith(&ogUrl, binding.JSON)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if len(ogUrl.Url) == 0 {
		utils.SendErrorResponse(c, http.StatusBadRequest, "invalid url")
	}

	urlMapsRepo := repo.UrlMapsRepo{DB: service.DB}
	var urlMap model.UrlMap
	urlMap.OriginalURL = ogUrl.Url

	var shortenedUrlPath string
	for {
		shortenedUrlPath = utils.GenerateRandomString(6)
		_, err = urlMapsRepo.FindUrlMap(shortenedUrlPath)
		if err != nil {
			break
		}
	}
	urlMap.ShortenedURLPath = shortenedUrlPath
	err = urlMapsRepo.InsertUrlMap(&urlMap)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	shortenedUrl := config.GetConfig().Server.BaseURL + "/" + shortenedUrlPath
	utils.SendSuccessResponse(c, http.StatusCreated, shortenedUrl)
}

func (service *URLMapService) RedirectToOriginalURL(c *gin.Context) {
	shortURLPath, found := c.Params.Get("shortenedURLPath")
	if found == false {
		utils.SendErrorResponse(c, http.StatusBadRequest, "missing urlMap_id in params")
	}
	//
	urlMapsRepo := repo.UrlMapsRepo{DB: service.DB}
	urlMap, err := urlMapsRepo.FindUrlMap(shortURLPath)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, err.Error())
	}
	if urlMap == nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Path Not Found!")
		return
	}

	c.Redirect(http.StatusFound, urlMap.OriginalURL)
}
