package pokecache

import (
	"net/http"
	"sync"
	"time"
)

type Cache struct{
	cache map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val http.Response
}


func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}


func (c Cache) IsEmpty() bool {
	return len(c.cache) == 0
}

func (c Cache) Add(key string, val http.Response) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[key]
	if !ok {
		c.cache[key] = cacheEntry{
			createdAt: time.Now(),
			val: val,
		}
	}

}


func (c Cache) Get(key string) (http.Response, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	if ok {
		return val.val, true
	}else {
		return http.Response{}, false
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cache {
			if time.Since(v.createdAt) > interval {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}