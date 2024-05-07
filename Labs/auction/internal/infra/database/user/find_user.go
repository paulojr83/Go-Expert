package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/logger"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/user_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id     string `bson:"_id"`
	UserId string `bson:"user_id"`
	Name   string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{Collection: database.Collection("users")}
}

func (ur *UserRepository) FindUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"user_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this id =%s", userId), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("User not found with this id =%s", userId))
		}

		logger.Error("Error trying to find user by Id", err)
		return nil, internal_error.NewInternalServerError("Error trying to find user by Id")
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.UserId,
		Name: userEntityMongo.Name,
	}
	return userEntity, nil
}
