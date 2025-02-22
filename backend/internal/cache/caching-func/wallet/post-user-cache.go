package walletCachingfunc

import (
	"context"
	"encoding/json"
	"time"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/wallet"
)

func CacheWallet(wallet database.Wallet) error {

	cacheExpiration := 7 * 24 * time.Hour // 7 days

	ctx := context.Background()

	RedisClient, _ := cache.RedisInit()
	walletMap := map[string]interface{}{
		"id":        wallet.Id,
		"UserId":    wallet.UserId,
		"Name":      wallet.Name,
		"createdAt": wallet.CreatedAt,
		"updatedAt": wallet.UpdatedAt,
	}

	walletDataJSON, err := json.Marshal(walletMap)
	if err != nil {
		return err
	}

	redisKey := "wallet:" + wallet.Id
	err = RedisClient.Set(ctx, redisKey, walletDataJSON, cacheExpiration).Err()
	if err != nil {
		return err
	}

	return nil

}
