package memcached_helpers

func (m *MemcacheCache) Close() {
	defer m.client.Close()
}
