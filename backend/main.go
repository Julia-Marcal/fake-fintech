package main

import (
	"fmt"

	services "github.com/Julia-Marcal/fake-fintech/config"
	db "github.com/Julia-Marcal/fake-fintech/internal/database"
	router "github.com/Julia-Marcal/fake-fintech/internal/http/router"
)

func main() {
	database_conn := db.NewPostgres()

	sqlDB, err := database_conn.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get database handle: %v", err))
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("database not reachable: %v", err))
	}

	services.NewDB()
	router.RunServer()
}
