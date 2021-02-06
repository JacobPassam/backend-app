package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	Database string = "backenddb"
	UsersCollection string = "users"
	PostsCollection string = "posts"
)

type MongoHandler struct {
	Client *mongo.Client
}

func NewMongoHandler() MongoHandler {

	host := "localhost:27017"
	username := "admin1"
	password := "123"

	mongoOptions := options.Client()

	mongoOptions.SetAuth(options.Credential{Username: username, Password: password})
	mongoOptions.ApplyURI("mongodb://" + host)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	client, err := mongo.Connect(ctx, mongoOptions)

	if err != nil {
		panic("Unable to connect to MongoDB: " + err.Error())
	}

	fmt.Println("Successfully connected to MongoDB")
	return MongoHandler{Client: client}

}

func (mongo *MongoHandler) Disconnect() {
	ctx := context.Background()

	mongo.Client.Disconnect(ctx)
	fmt.Println("Disconnected from MongoDB")
}
