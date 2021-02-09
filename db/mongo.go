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

var handler *MongoHandler

type MongoHandler struct {
	Client *mongo.Client
}

func NewMongoHandler() {

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

	handler = &MongoHandler{Client: client}
}

func Disconnect() error {
	ctx := context.Background()

	if err := handler.Client.Disconnect(ctx); err != nil {
		return err
	}

	fmt.Println("Disconnected from MongoDB")
	return nil
}

func Get() *MongoHandler {
	return handler
}

func GetCollection(name string) *mongo.Collection {
	cfg := config.Get()
	return handler.Client.Database(cfg.Mongo.Database).Collection(name)
}