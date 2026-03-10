package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/olucascdev/crud-user-golang/src/configuration/database/mongodb"
	"github.com/olucascdev/crud-user-golang/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Initialize database
	ctx := context.Background()
	database, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Init dependencies
	userController, _ := initDependencies(database)
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
