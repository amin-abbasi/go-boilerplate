package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB     *mongo.Database
	uri    = "mongodb://localhost:27017"
	dbName = "go_db_test"
	client *mongo.Client
)

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
			log.Fatalf(">>> Error connecting to MongoDB: %v\n", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
			log.Fatalf(">>> Error pinging MongoDB: %v\n", err)
	}

	// Set the database instance
	DB = client.Database(dbName)
	log.Println("<<< Connected to MongoDB >>>")
}

func DisconnectDB() {
	if client != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := client.Disconnect(ctx); err != nil {
					log.Fatalf(">>> Error disconnecting from MongoDB: %v\n", err)
			}
			log.Println("<<< Disconnected from MongoDB >>>")
	}
}