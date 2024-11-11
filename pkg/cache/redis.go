package cache

import (
	"app/config"
	"github.com/go-redis/redis"
)

type Cache struct {
	*redis.Client
}

func NewRedis(cfg *config.Config) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	return &Cache{client}
}
