package service

import (
	"github.com/olucascdev/crud-user-golang/src/controller/rest_err"
	"github.com/olucascdev/crud-user-golang/src/model"
)

func (*userDomainService) UpdateUser(
	userId string, userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
