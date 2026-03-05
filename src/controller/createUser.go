package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/olucascdev/crud-user-golang/src/controller/model/request"
	"github.com/olucascdev/crud-user-golang/src/controller/rest_err"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds,  error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return

	}
	fmt.Println(userRequest)
}
