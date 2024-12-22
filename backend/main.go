package main

import (
	db "github.com/Julia-Marcal/reusable-api/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var client *mongo.Client = db.NewMongoDB()
	_ = db.GetCollections(client)

	// router.StartRouter()
}
