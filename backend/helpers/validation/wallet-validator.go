package validation

import (
	wallet "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
	"github.com/go-playground/validator/v10"
)

type WalletStruct struct {
	Id     string `validate:"required,uuid4"`
	UserId string `validate:"required"`
}

func WalletValidator(w wallet.Wallet) bool {
	validate := validator.New()

	walletStruct := WalletStruct{
		Id:     w.Id,
		UserId: w.UserId,
	}

	err := validate.Struct(walletStruct)

	errors := ErrorHandler(err)

	return errors
}
