package redis_helpers

import "fmt"

func (r *RedisCache) Get(key string) ([]byte, error) {
	item, err := r.client.Get(key).Result()

	if err != nil {
		return nil, err
	}

	fmt.Println("Cache hit")

	return []byte(item), nil
}
