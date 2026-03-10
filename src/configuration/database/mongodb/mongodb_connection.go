package mongodb

import (
	"context"
	"log"
	"os"

	"github.com/olucascdev/crud-user-golang/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var ()

func InitConnection() {
	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("Error connecting to MongoDB")
		os.Exit(1)
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	logger.Info("Connected to MongoDB")
}
