package wallet_acoes_service

import (
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/walletAcoes/queries"
)

type WalletAcoesService interface {
	FindAcoesByWallet(walletId string) (interface{}, error)
}

type walletAcoesServiceImpl struct{}

func NewWalletAcoesService() WalletAcoesService {
	return &walletAcoesServiceImpl{}
}

func (s *walletAcoesServiceImpl) FindAcoesByWallet(walletId string) (interface{}, error) {
	return queries.FindAcoesByWallet(walletId)
}
