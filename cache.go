package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	Data      interface{}
	Timestamp time.Time
}

type Cache struct {
	items sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

// Set adds a value to the cache
// The key is the request URI
// The value is the response status code
// The cache is thread-safe
func (c *Cache) Set(key string, value interface{}) {
	c.items.Store(key, CacheItem{
		Data:      value,
		Timestamp: time.Now(),
	})
}

// Get retrieves a value from the cache
// If the value is not found, the second return value is false
// If the value is found, the second return value is true
// If the value is found but has expired, it is deleted from the cache
// and the second return value is false
// The cache expires after 30 seconds
// The cache is implemented using a sync.Map
// The key is the request URI
// The value is the response status code
// The cache is thread-safe
func (c *Cache) Get(key string) (interface{}, bool) {
	item, found := c.items.Load(key)
	if !found {
		return nil, false
	}
	fmt.Println("Cache Get!")
	fmt.Println("key:", key)
	fmt.Println("item:", item)

	cachedItem := item.(CacheItem)
	if time.Since(cachedItem.Timestamp) > 30*time.Second {
		c.items.Delete(key)
		return nil, false
	}

	return cachedItem.Data, true
}

func (c *Cache) Delete(key string) {
	c.items.Delete(key)
}
