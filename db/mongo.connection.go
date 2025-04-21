package db

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConnection interface {
	Connection() *mongo.Client
}

type mongoConnection struct{}

func NewMongoConnection() MongoConnection {
	return &mongoConnection{}
}
func (*mongoConnection) Connection() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Error with the mongo uro environment")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
