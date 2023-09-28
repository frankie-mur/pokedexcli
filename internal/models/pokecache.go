package models

import (
	"fmt"
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
	fmt.Println("adding")
	c.mutex.Lock()
	fmt.Printf("Adding to cache: %s\n", key)
	c.entry[key] = &cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Printf("Getting cache key %s\n", key)
	c.mutex.Lock()
	obj, ok := c.entry[key]
	if !ok {
		fmt.Println("Did not find entry")
		c.mutex.Unlock()
		return []byte{}, false
	}
	fmt.Printf("Found entry\n")
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
