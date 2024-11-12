package cache

import (
	"app/config"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Cache struct {
	*redis.Client
}

func NewRedis(ctx context.Context, cfg *config.Config) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	return &Cache{client}
}

// LockMultipleKeys try multiple lock keys
func (c *Cache) LockMultipleKeys(ctx context.Context, keys []string, timeout time.Duration) bool {
	err := c.Watch(ctx, func(tx *redis.Tx) error {
		for _, key := range keys {
			pipe := tx.Pipeline()
			pipe.SetNX(ctx, key, key, timeout)
			_, err := pipe.Exec(ctx)
			if err != nil {
				log.Fatalf("Failed to lock keys: %v", err)
				return err
			}
		}
		return nil
	}, keys...)
	if err != nil {
		log.Fatalf("Failed to lock keys: %v", err)
		return false
	}
	return true
}

// ReleaseMultipleLock release lock
func (c *Cache) ReleaseMultipleLock(ctx context.Context, keys []string) bool {
	err := c.Watch(ctx, func(tx *redis.Tx) error {
		err := c.Del(ctx, keys...).Err()
		if err != nil {
			log.Printf("Failed to release lock: %v", err)
		}
		return nil
	}, keys...)
	if err != nil {
		log.Fatalf("Failed to lock keys: %v", err)
		return false
	}
	return true
}
