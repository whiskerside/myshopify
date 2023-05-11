package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/whiskerside/myshopify/conf"
)

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", conf.Env.Redis.Host, conf.Env.Redis.Port),
		DB:          conf.Env.Redis.Db,
		PoolSize:    conf.Env.Redis.PoolSize,
		MaxRetries:  3,
		IdleTimeout: 10 * time.Second,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err == redis.Nil || err != nil {
		panic(fmt.Sprintf("Redis init failed, err: %+v\n", err))
	}
	logs.Info().Msg(fmt.Sprintf("Redis init... %+v", pong))
	return rdb
}
