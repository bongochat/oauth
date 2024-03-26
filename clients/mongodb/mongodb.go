package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_URL := os.Getenv("MONGODB_URL")

	// Set up MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Update the URI with your MongoDB connection string
	uri := DB_URL
	clientOptions := options.Client().ApplyURI(uri)
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Client = c
}

func GetCollections() *mongo.Collection {
	collection := Client.Database("oauth_db").Collection("tokens")
	// Create a unique index on the "username" field
	index := mongo.IndexModel{
		Keys:    bson.M{"accesstoken": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.Fatal(err)
	}
	return collection
}
