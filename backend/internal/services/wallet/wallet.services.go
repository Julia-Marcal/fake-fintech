package service

import (
	"errors"
	"fmt"

	"github.com/Julia-Marcal/fake-fintech/helpers/validation"
	cache "github.com/Julia-Marcal/fake-fintech/internal/cache/caching-func/wallet"
	"github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
	queries "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet/queries"
)

func GetAllWalletsService(userId string) ([]wallet.Wallet, error) {
	return queries.FindWallets(userId)
}

func CreateWalletService(walletData wallet.Wallet) (*wallet.Wallet, error) {
	if !validation.WalletValidator(walletData) {
		return nil, errors.New("Invalid input data")
	}

	if err := cache.CacheWallet(walletData); err != nil {
		return nil, errors.New("Failed to cache wallet")
	}

	if err := queries.Create(&walletData); err != nil {
		return nil, errors.New("Failed to create wallet")
	}

	return &walletData, nil
}

func GetWalletService(walletId string) (*wallet.Wallet, error) {
	cachedWallet, cacheErr := cache.GetCachedWallet(walletId)
	if cacheErr != nil {
		fmt.Println("Cache retrieval error:", cacheErr)
	}

	if cachedWallet.Id != "" {
		return &cachedWallet, nil
	}

	walletData, err := queries.FindWallet(walletId)
	if err != nil {
		return nil, err
	}

	if walletData == nil {
		return nil, nil
	}

	if err := cache.CacheWallet(*walletData); err != nil {
		return nil, errors.New("Failed to cache wallet")
	}

	return walletData, nil
}
