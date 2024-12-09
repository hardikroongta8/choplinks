package repository

import (
	"context"
	"github.com/hardikroongta8/choplinks/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	UserCollection *mongo.Collection
}

func (repo *UserRepo) InsertUser(c context.Context, user *model.User) (interface{}, error) {
	res, err := repo.UserCollection.InsertOne(c, user)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (repo *UserRepo) FindUserByID(c context.Context, userId string) (*model.User, error) {
	var user model.User
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	err = repo.UserCollection.FindOne(
		c, bson.D{{Key: "_id", Value: objectId}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) UpdateUserByID(c context.Context, userId string, up *model.User) (int64, error) {
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return 0, err
	}
	res, err := repo.UserCollection.UpdateOne(
		c, bson.D{{Key: "_id", Value: objectId}},
		bson.D{{Key: "$set", Value: up}},
	)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}
