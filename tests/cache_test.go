package tests

import (
	"context"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	ctx := context.Background()
	_, cache := NewDBAndCache(t)
	keys := []string{"user1", "user2", "user3"}
	defer func() {
		cache.ReleaseMultipleLock(ctx, keys)
	}()
	lock := cache.LockMultipleKeys(ctx, keys, 30*time.Second)
	if !lock {
		t.Log("lock success status:", lock)
	}
}
