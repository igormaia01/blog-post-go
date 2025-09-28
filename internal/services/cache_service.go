package services

import (
	"sync"
	"time"
)

// CacheItem represents a cached item with expiration
type CacheItem struct {
	Value     interface{}
	ExpiresAt time.Time
}

// CacheService provides caching functionality
type CacheService interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
	Delete(key string)
	Clear()
}

// MemoryCache implements an in-memory cache
type MemoryCache struct {
	items map[string]CacheItem
	mutex sync.RWMutex
}

// NewMemoryCache creates a new memory cache instance
func NewMemoryCache() *MemoryCache {
	cache := &MemoryCache{
		items: make(map[string]CacheItem),
	}
	
	// Start cleanup goroutine
	go cache.cleanup()
	
	return cache
}

// Get retrieves a value from the cache
func (mc *MemoryCache) Get(key string) (interface{}, bool) {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()
	
	item, exists := mc.items[key]
	if !exists {
		return nil, false
	}
	
	// Check if expired
	if time.Now().After(item.ExpiresAt) {
		return nil, false
	}
	
	return item.Value, true
}

// Set stores a value in the cache with TTL
func (mc *MemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	
	mc.items[key] = CacheItem{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// Delete removes a value from the cache
func (mc *MemoryCache) Delete(key string) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	
	delete(mc.items, key)
}

// Clear removes all items from the cache
func (mc *MemoryCache) Clear() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	
	mc.items = make(map[string]CacheItem)
}

// cleanup removes expired items from the cache
func (mc *MemoryCache) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		mc.mutex.Lock()
		now := time.Now()
		for key, item := range mc.items {
			if now.After(item.ExpiresAt) {
				delete(mc.items, key)
			}
		}
		mc.mutex.Unlock()
	}
}

// NoOpCache is a no-operation cache for testing or when caching is disabled
type NoOpCache struct{}

// NewNoOpCache creates a new no-op cache instance
func NewNoOpCache() *NoOpCache {
	return &NoOpCache{}
}

// Get always returns false (not found)
func (noc *NoOpCache) Get(key string) (interface{}, bool) {
	return nil, false
}

// Set does nothing
func (noc *NoOpCache) Set(key string, value interface{}, ttl time.Duration) {
	// No-op
}

// Delete does nothing
func (noc *NoOpCache) Delete(key string) {
	// No-op
}

// Clear does nothing
func (noc *NoOpCache) Clear() {
	// No-op
}
