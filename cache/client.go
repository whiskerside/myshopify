package cache

import (
	"github.com/go-redis/cache/v8"
	"github.com/whiskerside/myshopify/storage"
)

func Client() *cache.Cache {
	cacheClient := cache.New(&cache.Options{
		Redis: storage.RedisClient(),
	})
	return cacheClient
}
