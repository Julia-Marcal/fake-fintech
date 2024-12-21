package main

import (
	db "github.com/Julia-Marcal/reusable-api/internal/database"
)

func main() {
	db.NewMongoDB()
	// router.StartRouter()
}
