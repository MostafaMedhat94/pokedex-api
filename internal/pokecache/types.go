package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
	ttl     time.Duration
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}
