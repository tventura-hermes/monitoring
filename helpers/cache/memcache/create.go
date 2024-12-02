package memcached_helpers

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func (m *MemcacheCache) Set(key string, value []byte, expiration int) error {
	fmt.Println("saving cache")

	if err := m.client.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(expiration),
	}); err != nil {
		return fmt.Errorf("error setting cache: %v", err)
	}

	return nil
}
