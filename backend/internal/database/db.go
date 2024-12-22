package repository

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Julia-Marcal/reusable-api/config/env"
	database "github.com/Julia-Marcal/reusable-api/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	client *mongo.Client
)

type Collections struct {
	Users *mongo.Collection
}

func NewMongoDB() *mongo.Client {
	once.Do(func() {
		connectionStr := env.GetMongoConnectionString()
		client, err := mongo.NewClient(options.Client().ApplyURI(connectionStr))
		if err != nil {
			panic(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Successfully connected to MongoDB")
	})

	return client
}

func GetCollections(client *mongo.Client) *Collections {
	return &Collections{
		Users: database.CreateUsersCollection(client),
	}
}
