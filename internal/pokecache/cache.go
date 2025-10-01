package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println(c.entries)
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: value}
}

func (c *Cache) Get(key string) (value []byte, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, isMapContainsKey := c.entries[key]; isMapContainsKey {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			entryDuration := time.Since(entry.createdAt)
			if entryDuration > c.interval {
				delete(c.entries, key)

			}
		}
		c.mu.Unlock()
	}

}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{entries: make(map[string]cacheEntry), interval: interval}
	go newCache.reapLoop()
	return newCache
}
