package services

import (
	database "github.com/Julia-Marcal/fake-fintech/internal/user"
	queries "github.com/Julia-Marcal/fake-fintech/internal/user/queries"
)

func NewDB() {
	user := &database.User{
		Name:     "Julia",
		LastName: "Marcal",
		Age:      18,
		Email:    "gmail@gmail.com",
		Password: "123456",
	}

	queries.Create(user)
}
