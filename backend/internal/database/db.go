package repository

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Julia-Marcal/reusable-api/config/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	client *mongo.Client
	db     *mongo.Database
)

// NewMongoDB initializes and returns the MongoDB database instance.
func NewMongoDB() *mongo.Database {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		connectionStr := env.GetMongoConnectionString()
		fmt.Println("Connecting to MongoDB with connection string:", connectionStr)

		clientOptions := options.Client().ApplyURI(connectionStr)
		var err error
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			panic(fmt.Sprintf("Failed to ping MongoDB: %v", err))
		}

		databaseName := os.Getenv("MONGO_DATABASE")
		db = client.Database(databaseName)
		fmt.Println("Successfully connected to MongoDB:", databaseName)

		initializeCollections(db)
	})

	return db
}

func initializeCollections(db *mongo.Database) {
	usersCollection := db.Collection("users")

	indexModel := mongo.IndexModel{
		Keys: map[string]interface{}{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := usersCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		fmt.Printf("Failed to create index for users collection: %v\n", err)
	} else {
		fmt.Println("Successfully created index for users collection")
	}
}
