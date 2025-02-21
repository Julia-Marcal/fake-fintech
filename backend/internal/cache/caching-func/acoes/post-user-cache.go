package acoesCachingfunc

import (
	"context"
	"encoding/json"
	"time"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/acoes"
)

func CacheAcoes(acoes database.Acoes) error {

	cacheExpiration := 7 * 24 * time.Hour // 7 days

	ctx := context.Background()

	RedisClient, _ := cache.RedisInit()
	acoesMap := map[string]interface{}{
		"id":        acoes.Id,
		"Name":      acoes.Name,
		"Type":      acoes.Type,
		"Price":     acoes.Price,
		"Quantity":  acoes.Quantity,
		"createdAt": acoes.CreatedAt,
		"updatedAt": acoes.UpdatedAt,
	}

	acoesDataJSON, err := json.Marshal(acoesMap)
	if err != nil {
		return err
	}

	redisKey := "acoes:" + acoes.Id
	err = RedisClient.Set(ctx, redisKey, acoesDataJSON, cacheExpiration).Err()
	if err != nil {
		return err
	}

	return nil

}
