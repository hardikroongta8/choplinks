package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var dbClient *mongo.Client = nil

func Connect(dbUri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("Database connection ping failed: %v", err)
	}
	log.Println("Successfully connected to the database!")
	dbClient = client
	return dbClient
}

func GetDatabase(dbName string) *mongo.Database {
	if dbClient == nil {
		log.Fatal("Database client is not initialized. Call Connect first.")
	}
	return dbClient.Database(dbName)
}

func Disconnect() {
	if dbClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		if err := dbClient.Disconnect(ctx); err != nil {
			log.Fatalf("Error while disconnecting from the database: %v", err)
		}
		log.Println("Disconnected from the database!")
	}
}
