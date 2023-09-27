package models

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry map[string]*cacheEntry
	mutex sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{entry: make(map[string]*cacheEntry)}
	go cache.readLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.entry[key] = &cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	obj, ok := c.entry[key]
	if !ok {
		return []byte{}, false
	}
	c.mutex.Unlock()
	return obj.val, true
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, item := range c.entry {
			if time.Since(item.createdAt) > interval {
				delete(c.entry, key)
			}
		}
		c.mutex.Unlock()
	}
}
