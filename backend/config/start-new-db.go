package services

import (
	"fmt"

	database "github.com/Julia-Marcal/reusable-api/internal/user"
	queries "github.com/Julia-Marcal/reusable-api/internal/user/queries"
)

func NewDB() {
	user := &database.User{
		Name:     "Julia",
		LastName: "Marcal",
		Age:      18,
		Email:    "gmail@gmail.com",
		Password: "123456",
	}

	user_inf := queries.Create(user)
	fmt.Println(user_inf)
}
