package pokecache

import (
	"sync"
	"time"
)
type Cache interface {
	Add(key string, val []byte)
	Get(key string) (val []byte, check bool)
	reapLoop(interval time.Duration)
} 
type pokecache struct{
	c map[string]cacheEntry
	mu *sync.Mutex
}
type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func (c *pokecache) Add(key string, val []byte) {

	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}

	c.c[key] = entry 
}

func (c *pokecache) Get(key string) (val []byte, check bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.c[key]

	if !ok {
		return nil, false 
	}

	return entry.val, true
}

func (c *pokecache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.c {
			if now.Sub(entry.createdAt) > interval {
				delete(c.c, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) Cache{
	pcache := &pokecache{
		c:  map[string]cacheEntry{},
		mu: &sync.Mutex{},
	}

	go pcache.reapLoop(interval) // run cleanup in background

	return pcache
}


