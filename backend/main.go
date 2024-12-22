package main

import (
	db "github.com/Julia-Marcal/reusable-api/internal/database"
)

func main() {
	client := db.NewMongoDB()
	_ = db.GetCollections(client)

	// router.StartRouter()
}
