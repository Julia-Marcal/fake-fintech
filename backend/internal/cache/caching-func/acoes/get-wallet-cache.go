package acoesCachingfunc

import (
	"context"
	"fmt"
	"time"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
)

func GetCachedAcoes(id string) (database.Acoes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	RedisClient, _ := cache.RedisInit()

	walletData, err := RedisClient.HGetAll(ctx, "acoes:"+id).Result()
	if err != nil {
		return database.Acoes{}, err
	}

	if len(walletData) == 0 {
		return database.Acoes{}, fmt.Errorf("investment allocation not found in cache")
	}

	var wallet database.Acoes

	return wallet, nil
}
