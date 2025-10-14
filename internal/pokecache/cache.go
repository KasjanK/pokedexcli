package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt 	time.Time
	val			[]byte
}

type Cache struct {
	entryList	map[string]cacheEntry
	mutex		*sync.Mutex
}

func NewCache(interval time.Duration) Cache{
	cache := Cache{
		entryList: 	make(map[string]cacheEntry),
		mutex: 		&sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.entryList[key] = cacheEntry{
		createdAt: 	time.Now(), 
		val: 		val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	entry, ok := cache.entryList[key]
	return entry.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)	
	for range ticker.C {
		cache.reap(time.Now(), interval)
	}
}

func (cache *Cache) reap(now time.Time, last time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for k, v := range cache.entryList {
		if v.createdAt.Before(now.Add(-last)) {
			delete(cache.entryList, k)
		}
	}
}
