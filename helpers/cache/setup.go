package cache_helpers

import (
	"context"
	memcached_helpers "demo/helpers/cache/memcache"
	redis_helpers "demo/helpers/cache/redis"
	"errors"
	"fmt"
	"os"
)

type Cache interface {
	Set(key string, value []byte, expiration int) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close()
}

func NewCache(ctx context.Context) (Cache, error) {
	cacheType := os.Getenv("CACHE_TYPE")
	switch cacheType {
	case "redis":
		fmt.Println("using redis")
		return redis_helpers.SetupRedisClient(ctx), nil
	case "memcached":
		fmt.Println("using memcache")
		return memcached_helpers.SetupMemcachedClient(), nil
	case "nocache":
		return nil, nil
	default:
		return nil, errors.New("unsupported cache backend")
	}
}
