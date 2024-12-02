package memcached_helpers

func (m *MemcacheCache) Delete(key string) error {
	return m.client.Delete(key)
}
