package memcached_helpers

import (
	"fmt"
)

func (m *MemcacheCache) Get(key string) ([]byte, error) {

	item, err := m.client.Get(key)

	fmt.Println("get cache")
	if err != nil {
		return nil, err
	}

	fmt.Println("Cache hit")

	return item.Value, nil
}
