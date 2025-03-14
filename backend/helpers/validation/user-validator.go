package validation

import (
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"

	"github.com/go-playground/validator/v10"
)

type UserStruct struct {
	Name     string `validate:"required"`
	LastName string `validate:"required"`
	Age      string `validate:"gte=0, lte=130"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string `validate:"required,oneof=admin user"`
}

func UserValidator(user database.User) bool {
	validate := validator.New()
	err := validate.Struct(user)

	errors := ErrorHandler(err)

	return errors
}
