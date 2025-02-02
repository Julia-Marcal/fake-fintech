package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(err any) bool {
	var errors bool

	if err != nil {
		fmt.Println("Validation failed:")
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("Field: %s, Error: %s\n", e.Field(), e.Tag())
		}
		errors = true
	} else {
		fmt.Println("Validation succeeded")
		errors = false
	}

	return errors
}
