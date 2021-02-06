package db

import (
	"backend-app/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

const (
	UsersCollection string = "users"
	PostsCollection string = "posts"
)

type MongoHandler struct {
	Client *mongo.Client
}

func NewMongoHandler() MongoHandler {

	cfg := config.Get()
	host := cfg.Mongo.Host
	port := cfg.Mongo.Port
	username := cfg.Mongo.Username
	password := cfg.Mongo.Password
	database := cfg.Mongo.Database

	mongoOptions := options.Client()

	mongoOptions.SetAuth(options.Credential{Username: username, Password: password, AuthSource: database})
	mongoOptions.ApplyURI("mongodb://" + host + ":" + strconv.Itoa(port))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, mongoOptions)

	if err != nil {
		panic("Unable to connect to MongoDB: " + err.Error())
	}

	fmt.Println("Successfully connected to MongoDB")
	return MongoHandler{Client: client}

}

func (mongo *MongoHandler) Disconnect() error {
	ctx := context.Background()

	if err := mongo.Client.Disconnect(ctx); err != nil {
		return err
	}

	fmt.Println("Disconnected from MongoDB")
	return nil
}
