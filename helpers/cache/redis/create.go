package redis_helpers

import (
	"fmt"
	"time"
)

func (r *RedisCache) Set(key string, value []byte, expiration int) error {
	fmt.Println("Creating cache")
	return r.client.Set(key, value, time.Duration(expiration)*time.Second).Err()
}
