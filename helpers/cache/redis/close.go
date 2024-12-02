package redis_helpers

func (r *RedisCache) Close() {
	defer r.client.Close()
}
