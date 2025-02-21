package validation

import (
	acoes "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
	"github.com/go-playground/validator/v10"
)

type AcoesStructure struct {
	Name     string  `validate:"required"`
	Type     string  `validate:"required"`
	Price    float64 `validate:"required"`
	Quantity float64 `validate:"required"`
}

func AcoesWalletValidator(a acoes.Acoes) bool {
	validate := validator.New()

	acoesStruct := AcoesStructure{
		Name:     a.Name,
		Type:     a.Type,
		Price:    a.Price,
		Quantity: a.Quantity,
	}

	err := validate.Struct(acoesStruct)
	return err == nil
}
