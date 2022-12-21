package db

import (
	"context"
	"log"
	"time"

	"github.com/hu553in/gist-clone/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func Init() {
	mongoUrl := config.GetConfig().GetString("db.url")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal("Failed to create MongoDB client", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}

	log.Println("Connected to MongoDB")
	dbClient = client
}

func GetClient() *mongo.Client {
	return dbClient
}
