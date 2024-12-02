package redis_helpers

import (
	"context"
	"os"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func SetupRedisClient(ctx context.Context) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})
	return &RedisCache{
		client: client,
		ctx:    ctx,
	}
}
