package redis_helpers

func (r *RedisCache) Delete(key string) error {
	return r.client.Del(key).Err()
}
