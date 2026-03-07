package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olucascdev/crud-user-golang/src/configuration/logger"
	"github.com/olucascdev/crud-user-golang/src/configuration/validation"
	"github.com/olucascdev/crud-user-golang/src/controller/model/request"
	"github.com/olucascdev/crud-user-golang/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error validating user", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return

	}
	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "1",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusOK, response)

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)
}
