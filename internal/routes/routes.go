package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikroongta8/choplinks/internal/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func SetupRoutes(db *mongo.Database) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.GlobalMiddleware)
	AttachUserRouter(router.Group("/user"), db)

	protectedRouter := router.Group("/")
	protectedRouter.Use(middlewares.AuthMiddleware)
	protectedRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a protected route!")
	})

	return router
}
