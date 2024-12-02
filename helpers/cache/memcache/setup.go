package memcached_helpers

import (
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemcacheCache struct {
	client *memcache.Client
}

func SetupMemcachedClient() *MemcacheCache {
	client := memcache.New(os.Getenv("CACHE_HOST"))
	return &MemcacheCache{
		client: client,
	}
}
