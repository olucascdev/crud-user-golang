package service

import (
	"github.com/olucascdev/crud-user-golang/src/controller/rest_err"
	"github.com/olucascdev/crud-user-golang/src/model"
	"github.com/olucascdev/crud-user-golang/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
