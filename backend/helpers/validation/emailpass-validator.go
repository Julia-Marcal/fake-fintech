package validation

import (
	"github.com/go-playground/validator/v10"
)

type EmailPassStruct struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func EmailPassValidator(login any) bool {
	validate := validator.New()
	err := validate.Struct(login)

	errors := ErrorHandler(err)

	return errors
}
