package database

import (
	"context"
	"fmt"
	"time"

	env "github.com/Julia-Marcal/reusable-api/config/env"
	security "github.com/Julia-Marcal/reusable-api/helpers/security"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UUID      string             `bson:"uuid"`
	Name      string             `bson:"name"`
	LastName  string             `bson:"last_name"`
	Age       int32              `bson:"age"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func CreateUsersCollection(client *mongo.Client) *mongo.Collection {
	db := client.Database(env.GetDatabase())
	usersCollection := db.Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := usersCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		fmt.Printf("Failed to create index for users collection: %v\n", err)
	} else {
		fmt.Println("Successfully created index for users collection")
	}

	validationSchema := bson.M{
		"bsonType": "object",
		"required": []string{"uuid", "name", "last_name", "age", "email", "password", "created_at", "updated_at"},
		"properties": bson.M{
			"uuid":       bson.M{"bsonType": "string"},
			"name":       bson.M{"bsonType": "string"},
			"last_name":  bson.M{"bsonType": "string"},
			"age":        bson.M{"bsonType": "int"},
			"email":      bson.M{"bsonType": "string"},
			"password":   bson.M{"bsonType": "string"},
			"created_at": bson.M{"bsonType": "date"},
			"updated_at": bson.M{"bsonType": "date"},
		},
	}

	validator := bson.M{"$jsonSchema": validationSchema}
	opts := options.CreateCollection().SetValidator(validator)

	err = db.CreateCollection(ctx, "Users", opts)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			fmt.Println("Users collection already exists, skipping creation.")
		} else {
			fmt.Printf("Failed to create users collection: %v\n", err)
		}
	} else {
		fmt.Println("Successfully created users collection with validation rules")
	}

	return usersCollection
}

func (user *User) BeforeInsert() error {
	user.UUID = uuid.NewString()

	salt := env.SetSalt()
	_, hashedPassword, err := security.DeriveKey(user.Password, salt)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return nil
}
