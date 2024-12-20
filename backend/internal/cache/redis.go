package repository

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/redis/go-redis/v9"
)

func RedisInit() (*redis.Client, ratelimit.Store) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: client,
		Rate:        time.Second,
		Limit:       10,
	})

	return client, store
}
