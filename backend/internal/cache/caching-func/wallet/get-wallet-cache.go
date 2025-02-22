package walletCachingfunc

import (
	"context"
	"fmt"
	"time"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
)

func GetCachedWallet(id string) (database.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	RedisClient, _ := cache.RedisInit()

	walletData, err := RedisClient.HGetAll(ctx, "wallet:"+id).Result()
	if err != nil {
		return database.Wallet{}, err
	}

	if len(walletData) == 0 {
		return database.Wallet{}, fmt.Errorf("wallet not found in cache")
	}

	var wallet database.Wallet

	return wallet, nil
}
