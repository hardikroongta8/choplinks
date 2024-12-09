package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hardikroongta8/choplinks/internal/model"
	"github.com/hardikroongta8/choplinks/internal/repository"
	"github.com/hardikroongta8/choplinks/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserService struct {
	UserCollection *mongo.Collection
}

func (service *UserService) CreateUser(c *gin.Context) {
	c.Request.Context()
	var user model.User
	user.ShortenedLinks = []model.UrlMap{}
	err := c.ShouldBindBodyWith(&user, binding.JSON)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if len(user.Username) == 0 || len(user.Password) == 0 {
		utils.SendErrorResponse(c, http.StatusBadRequest, "invalid username or password")
	}

	repo := repository.UserRepo{UserCollection: service.UserCollection}
	insertID, err := repo.InsertUser(c.Request.Context(), &user)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	utils.SendSuccessResponse(c, http.StatusCreated, insertID)
}

func (service *UserService) FindUserByID(c *gin.Context) {
	userId, found := c.Params.Get("user_id")
	if found == false {
		utils.SendErrorResponse(c, http.StatusBadRequest, "missing user_id in params")
	}

	repo := repository.UserRepo{UserCollection: service.UserCollection}
	user, err := repo.FindUserByID(c.Request.Context(), userId)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	utils.SendSuccessResponse(c, http.StatusOK, user)
}

func (service *UserService) UpdateUserByID(c *gin.Context) {
	userId, found := c.Params.Get("user_id")
	if found == false {
		utils.SendErrorResponse(c, http.StatusBadRequest, "missing user_id in params")
	}

	if userId == "" {
		utils.SendErrorResponse(c, http.StatusBadRequest, "invalid user id")
	}

	var updates model.User
	err := c.ShouldBindBodyWith(&updates, binding.JSON)

	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "invalid request body")
	}

	repo := repository.UserRepo{UserCollection: service.UserCollection}
	count, err := repo.UpdateUserByID(c.Request.Context(), userId, &updates)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	utils.SendSuccessResponse(c, http.StatusOK, count)
}
