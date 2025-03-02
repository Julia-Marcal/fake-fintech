package userCachingfunc

import (
	"context"
	"fmt"
	"time"

	cache "github.com/Julia-Marcal/fake-fintech/internal/cache"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"
)

func GetCachedUser(id string) (database.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	RedisClient, _ := cache.RedisInit()

	userData, err := RedisClient.HGetAll(ctx, "user:"+id).Result()
	if err != nil {
		return database.User{}, err
	}

	if len(userData) == 0 {
		return database.User{}, fmt.Errorf("user not found in cache")
	}

	var user database.User

	return user, nil
}
