package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type EmailPassStruct struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func EmailPassValidator(login any) bool {

	loginData, ok := login.(EmailPassStruct)
	if !ok {
		fmt.Println("Invalid type passed to EmailPassValidator")
		return false
	}

	validate := validator.New()
	err := validate.Struct(loginData)
	errors := ErrorHandler(err)

	return errors
}
