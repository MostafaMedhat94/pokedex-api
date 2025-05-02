package pokecache

import (
	"sync"
	"time"
)

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
		ttl:     ttl,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()

	defer c.mu.Unlock()

	entry, ok := c.entries[key]

	if !ok {
		return nil, false
	}

	if time.Since(entry.createdAt) > c.ttl {
		delete(c.entries, key)
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > c.ttl {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
