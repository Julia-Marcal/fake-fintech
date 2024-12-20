package main

import (
	db "github.com/Julia-Marcal/reusable-api/internal/database"
	router "github.com/Julia-Marcal/reusable-api/internal/http/router"
)

func main() {
	db.NewPostgres()
	router.StartRouter()
}
