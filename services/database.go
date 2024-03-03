package services

import (
	"context"
	"log"
	"time"

	"github.com/amin-abbasi/go-boilerplate/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB     *mongo.Database
	client *mongo.Client
)

func tryConnect() (*mongo.Client, error) {
	// Initialize MongoDB connection parameters
	host := configs.GetEnvVariable("DB_HOST")
	port := configs.GetEnvVariable("DB_PORT")
	dbURL := "mongodb://" + host + ":" + port
	log.Printf(">>> DB URL: %v", dbURL)

	maxRetries := 5
	retryDelay := time.Second * 5

	// Retry connecting to MongoDB
	var err error
	for retries := 0; retries < maxRetries; retries++ {
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(dbURL))
		if err == nil {
			// Connection successful, exit retry loop
			return client, err
		}
		log.Printf("Error connecting to MongoDB: %v", err)
		log.Printf("Retrying in %v...", retryDelay)
		time.Sleep(retryDelay)
	}

	if err != nil {
		return nil, err
	}

	defer DisconnectDB()
	return nil, err
}

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TIME_OUT_DURATION)
	defer cancel()

	dbName := configs.GetEnvVariable("DB_NAME")

	// Connect to MongoDB
	var err error
	client, err = tryConnect()
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
		ctx, cancel := context.WithTimeout(context.Background(), configs.TIME_OUT_DURATION)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf(">>> Error disconnecting from MongoDB: %v\n", err)
		}
		log.Println("<<< Disconnected from MongoDB >>>")
	}
}
