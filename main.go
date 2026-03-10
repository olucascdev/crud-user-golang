package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/olucascdev/crud-user-golang/src/controller"
	"github.com/olucascdev/crud-user-golang/src/controller/routes"
	"github.com/olucascdev/crud-user-golang/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependencies

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
