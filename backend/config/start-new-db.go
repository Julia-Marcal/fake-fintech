package services

import (
	"fmt"

	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/user/queries"
)

func NewDB() {
	user := &database.User{
		Name:     "Julia",
		LastName: "Marcal",
		Age:      18,
		Email:    "gmail@gmail.com",
		Password: "123456",
		Role:     "admin",
	}

	err := queries.Create(user)
	if err != nil {
		fmt.Println("Error creating user:", err)
	}

}
