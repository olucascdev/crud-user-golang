package repository

import (
	"context"
	"os"

	"github.com/olucascdev/crud-user-golang/src/configuration/logger"
	"github.com/olucascdev/crud-user-golang/src/controller/rest_err"
	"github.com/olucascdev/crud-user-golang/src/model"
	"github.com/olucascdev/crud-user-golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := bson.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to updateUser",
			err,
			zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUser repository executed successfully",
		zap.String("userId", userId),
		zap.Int64("matchedCount", result.MatchedCount),
		zap.Int64("modifiedCount", result.ModifiedCount),
		zap.String("journey", "updateUser"))

	return nil

}
