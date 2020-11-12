package gotool

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var once sync.Once

type memItem struct {
	item                interface{}
	lastUpdatedTime     time.Time
	expiredIntervalSecs int64 //0 means never being expired, unit seconds
}

// MemCache is the simple memory cache
type MemCache struct {
	CurrentMap *ConcurrentMap
	Ctx        context.Context
	cancel     context.CancelFunc
}

// NewMemCache is to create a memory cache
func NewMemCache(ctx context.Context) *MemCache {
	c, cancel := context.WithCancel(ctx)
	var mc *MemCache
	once.Do(func() {
		mc = &MemCache{
			CurrentMap: CreateConcurrentMap(99),
			Ctx:        c,
			cancel:     cancel,
		}
		mc.startCleanExpiredItem()
	})
	return mc
}

// Put is to put the item to memory cache
// expiredIntervalSecs is the expired duration (seconds)
func (m *MemCache) Put(key string, item interface{}, expiredIntervalSecs int64) {
	mItem := &memItem{
		item:                item,
		expiredIntervalSecs: expiredIntervalSecs,
		lastUpdatedTime:     time.Now(),
	}
	m.CurrentMap.Set(StrKey(key), mItem)
}

// Get the item from the cache
// bool return value means: true = existing, false = not existing/expired
func (m *MemCache) Get(key string) (interface{}, bool) {

	mItem, ok := m.CurrentMap.Get(StrKey(key))
	if !ok {
		return nil, false
	}

	assertMItem, _ := mItem.(*memItem)

	if assertMItem.expiredIntervalSecs == 0 {
		return assertMItem.item, true
	}

	if time.Now().After(assertMItem.lastUpdatedTime.Add(time.Second * time.Duration(assertMItem.expiredIntervalSecs))) {
		return nil, false
	}
	return assertMItem.item, true
}

func (m *MemCache) Delete(key string) {
	m.CurrentMap.Del(StrKey(key))
}

func (m *MemCache) startCleanExpiredItem() {
	go func(ctx context.Context) {
		logrus.Info("start clean expired item")
		for {
			select {
			case <-ctx.Done():
				logrus.Info("clean expired item is cancelled.")
				return
			case <-time.After(time.Minute * 3):
				for _, item := range m.CurrentMap.partitions {
					mapV := item.m
					for k, v := range mapV {
						assertKey := k.(string)
						assertValue := v.(*memItem)
						// logrus.Infof("assertKey --> %v, assertValue --> %v", assertKey, assertValue)
						if time.Now().After(assertValue.lastUpdatedTime.Add(time.Second * time.Duration(assertValue.expiredIntervalSecs))) {
							m.Delete(assertKey)
						}
					}
				}
				logrus.Info("expired item has been clean")
			}
		}
	}(m.Ctx)
}

func (m *MemCache) Destroy() {
	m.cancel()
}
