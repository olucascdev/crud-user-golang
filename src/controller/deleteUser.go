package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olucascdev/crud-user-golang/src/configuration/logger"
	"github.com/olucascdev/crud-user-golang/src/controller/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser ontroller",
		zap.String("journey", "DeleteUser"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)

	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call DeleteUser service",
			err,
			zap.String("journey", "DeleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully", zap.String("userId", userId),
		zap.String("journey", "DeleteUser"))

	c.Status(http.StatusOK)
}
