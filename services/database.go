package services

import (
	"context"
	"log"

	"github.com/amin-abbasi/go-boilerplate/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB     *mongo.Database
	client *mongo.Client
)

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.TIME_OUT_DURATION)
	defer cancel()

	// Get Variables
	host := configs.GetEnvVariable("DB_HOST")
	port := configs.GetEnvVariable("DB_PORT")
	dbName := configs.GetEnvVariable("DB_NAME")
	dbURL := "mongodb://" + host + ":" + port

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
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
