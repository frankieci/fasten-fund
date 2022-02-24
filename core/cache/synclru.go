package cache

import "sync"

/**
 * Created by frankieci on 2022/1/1 4:28 pm
 */

// LRUCache is a thread-safe fixed size LRU cache.
type LRUCache struct {
	cache *lruCache
	lock  sync.RWMutex
}

// NewLRUCache creates an LRU of the given size.
func NewLRUCache(options ...Option) *LRUCache {
	return &LRUCache{cache: New(options...)}
}

// Add adds a value to the cache.
func (c *LRUCache) Add(key Key, value Value) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Add(key, value)
}

// Get looks up a key's value from the cache.
func (c *LRUCache) Get(key Key) (Value, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.cache.Get(key)
}

// Remove removes the provided key from the cache.
func (c *LRUCache) Remove(key Key) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Remove(key)
}

// RemoveOldest removes the oldest item from the cache.
func (c *LRUCache) RemoveOldest() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.RemoveOldest()
}

// Len returns the number of items in the cache.
func (c *LRUCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cache.Len()
}

// Clear purges all stored items from the cache.
func (c *LRUCache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cache.Clear()
}
