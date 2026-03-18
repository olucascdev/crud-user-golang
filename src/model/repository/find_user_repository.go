package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/olucascdev/crud-user-golang/src/configuration/logger"
	"github.com/olucascdev/crud-user-golang/src/controller/rest_err"
	"github.com/olucascdev/crud-user-golang/src/model"
	"github.com/olucascdev/crud-user-golang/src/model/repository/entity"
	"github.com/olucascdev/crud-user-golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository",
		zap.String("journey", "FindUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(&userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "FindUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("email", userEntity.Email),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "FindUserByEmail"))
	return converter.ConvertEntityToDomain(&userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID repository",
		zap.String("journey", "FindUserByID"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := entity.UserEntity{}

	objectId, _ := bson.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(&userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "FindUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByID repository executed successfully",
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "FindUserByID"))
	return converter.ConvertEntityToDomain(&userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword repository",
		zap.String("journey", "FindUserByEmailAndPassword"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(&userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByEmailAndPassword"))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "FindUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmailAndPassword repository executed successfully",
		zap.String("email", userEntity.Email),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "FindUserByEmailAndPassword"))
	return converter.ConvertEntityToDomain(&userEntity), nil
}
