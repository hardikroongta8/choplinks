package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikroongta8/choplinks/internal/middlewares"
	"github.com/hardikroongta8/choplinks/internal/service"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	urlMapService := service.URLMapService{DB: db}
	router := gin.Default()
	router.Use(middlewares.GlobalMiddleware)
	router.POST("/url/create", urlMapService.CreateUrlMap)

	router.GET("/:shortenedURLPath", urlMapService.RedirectToOriginalURL)

	return router
}
