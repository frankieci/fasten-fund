package cache

import "container/list"

/**
 * Created by frankieci on 2022/1/1 3:53 pm
 */

// lruCache is an LRU cache. It is not safe for concurrent access.
type lruCache struct {
	// MaxEntries is the maximum number of cache entries before
	// an item is evicted. Zero means no limit.
	MaxEntries int

	// OnEvicted optionally specifies a callback function to be
	// executed when an entry is purged from the cache.
	OnEvicted OnEvictedFunc

	list  *list.List
	cache map[Key]*list.Element
}

type Option func(*lruCache)
type OnEvictedFunc func(key Key, value Value)

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators
type Key = interface{}
type Value = interface{}

type entry struct {
	key   Key
	value Value
}

func newEntry(key Key, value Value) *entry {
	return &entry{key, value}
}

// New creates a new lruCache.
// If maxEntries is zero, the cache has no limit, and it's assumed
// that eviction is done by the caller.
func New(options ...Option) *lruCache {
	cache := &lruCache{list: list.New(), cache: make(map[Key]*list.Element)}

	for _, opt := range options {
		opt(cache)
	}
	return cache
}

func WithMaxEntries(maxEntries int) Option {
	return func(cache *lruCache) {
		cache.MaxEntries = maxEntries
	}
}

func WithOnEvicted(onEvicted OnEvictedFunc) Option {
	return func(cache *lruCache) {
		cache.OnEvicted = onEvicted
	}
}

// Add adds a value to the cache.
func (c *lruCache) Add(key Key, value Value) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.list = list.New()
	}

	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	elem := c.list.PushFront(newEntry(key, value))
	c.cache[key] = elem
	if c.MaxEntries != 0 && c.list.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// Get looks up a key's value from the cache.
func (c *lruCache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if elem, hit := c.cache[key]; hit {
		c.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return
}

// Remove removes the provided key from the cache.
func (c *lruCache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if elem, hit := c.cache[key]; hit {
		c.removeElement(elem)
	}
}

// RemoveOldest removes the oldest item from the cache.
func (c *lruCache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	elem := c.list.Back()
	if elem != nil {
		c.removeElement(elem)
	}
}

// Len returns the number of items in the cache.
func (c *lruCache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.list.Len()
}

// Clear purges all stored items from the cache.
func (c *lruCache) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.list = nil
	c.cache = nil
}

func (c *lruCache) removeElement(elem *list.Element) {
	c.list.Remove(elem)
	kv := elem.Value.(*entry)
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}
