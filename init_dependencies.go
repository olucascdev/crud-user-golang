package main

import (
	"github.com/olucascdev/crud-user-golang/src/controller"
	"github.com/olucascdev/crud-user-golang/src/model/repository"
	"github.com/olucascdev/crud-user-golang/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(database *mongo.Database) (
	controller.UserControllerInterface,
	error,
) {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service), nil
}
