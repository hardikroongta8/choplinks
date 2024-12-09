package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikroongta8/choplinks/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func AttachUserRouter(router *gin.RouterGroup, db *mongo.Database) {
	userCollection := db.Collection("users")
	userService := service.UserService{UserCollection: userCollection}
	router.POST("/create", userService.CreateUser)
	router.GET("/find/:user_id", userService.FindUserByID)
	router.PUT("/update/:user_id", userService.UpdateUserByID)
}
