package gotool

import (
	"context"
	"testing"
	"time"
)

func TestBasicOps(t *testing.T) {
	key := "TestKey"
	value := "Test"
	ctx := context.Background()
	cache := NewMemCache(ctx)
	cache.Put(key, value, 5)
	c, isExisting := cache.Get(key)
	if !isExisting {
		t.Error("failed to put item to cache")
	}
	str := c.(string)
	if str != value {
		t.Error("get the wrong value from cache")
	}
	time.Sleep(210 * time.Second)
	_, isExisting = cache.Get(key)
	if isExisting {
		t.Error("failed to set expired policy for cache")
	}
	cache.Destroy()
}