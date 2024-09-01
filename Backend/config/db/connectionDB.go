package db

import (
	"context"
	"log"
	"time"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var FreelancerCollection *mongo.Collection
var ClientCollection *mongo.Collection
var JobCollection *mongo.Collection
var ReviewsCollection *mongo.Collection

func ConnectDB(connectionString string) {

    clientOptions := options.Client().ApplyURI(connectionString)

    client, err := mongo.NewClient(clientOptions)

    if err != nil {
        log.Fatalf("Error creating MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    Client = client
    FreelancerCollection = client.Database("felagi").Collection("Freelancer")
    ClientCollection = client.Database("felagi").Collection("Client")
    JobCollection = client.Database("felagi").Collection("Job")
    ReviewsCollection = client.Database("felagi").Collection("Reviews")
}
