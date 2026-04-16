// Package cache provides a concurrent, TTL-based in-memory cache with automatic
// cleanup and statistics tracking. It is designed for high-concurrency scenarios
// where data needs to expire after a configured duration.
package cache

import (
	"fmt"
	"sync"
	"time"
)

// Item represents a cached entry with its value and expiration time.
type Item struct {
	value      interface{}
	expiration time.Time
}

// Cache is a thread-safe in-memory cache that automatically expires entries
// after a configured TTL (Time-To-Live). It supports concurrent access from
// multiple goroutines and provides statistics about cache performance.
type Cache struct {
	mu              sync.RWMutex
	items           map[string]Item
	ttl             time.Duration
	stats           Stats
	cleanupInterval time.Duration
	stopCleanup     chan struct{}
	stopped         bool
}

// Stats holds performance metrics for the cache, tracking hit/miss counts
// and the current cache size. All fields are safe for concurrent access.
type Stats struct {
	Hits int
	Miss int
	Size int
	mu   sync.RWMutex
}

// NewCache creates and initializes a new Cache instance with the specified
// TTL (Time-To-Live) for entries and cleanup interval for automatic removal
// of expired items.
//
// Parameters:
//   - ttl: Duration after which an entry expires and becomes eligible for removal
//   - cleanupInterval: How often the cache scans for and removes expired entries
//
// Returns:
//   - A pointer to the initialized Cache instance, ready for use
//
// The cleanup goroutine runs automatically and must be stopped with Stop()
// when the cache is no longer needed to prevent goroutine leaks.
func NewCache(ttl, cleanupInterval time.Duration) *Cache {
	c := &Cache{
		items:           make(map[string]Item),
		ttl:             ttl,
		cleanupInterval: cleanupInterval,
		stopCleanup:     make(chan struct{}),
		stopped:         false,
	}

	go c.startCleanup()
	return c
}

// Set stores a value in the cache with the configured TTL. If the key already
// exists, its value and expiration time are overwritten.
//
// Parameters:
//   - key: The identifier under which to store the value
//   - value: The data to cache (any type)
//
// This operation is thread-safe and automatically updates cache size statistics.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = Item{
		value:      value,
		expiration: time.Now().Add(c.ttl),
	}

	c.updateStatsSize()
}

// Get retrieves a value from the cache and automatically refreshes its TTL
// upon successful retrieval. Expired entries are removed and count as misses.
//
// Parameters:
//   - key: The identifier of the cached value to retrieve
//
// Returns:
//   - The cached value (any type) if found and not expired
//   - A boolean indicating whether the value was successfully retrieved
//
// The TTL refresh ensures frequently accessed entries remain in the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, exists := c.items[key]
	c.mu.RUnlock()

	if !exists {
		c.incrementMiss()
		return nil, false
	}

	now := time.Now()

	if now.After(item.expiration) {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		c.updateStatsSize()
		c.incrementMiss()
		return nil, false
	}

	c.mu.Lock()
	item.expiration = now.Add(c.ttl)
	c.items[key] = item
	c.mu.Unlock()

	c.incrementHit()
	return item.value, true
}

// Delete removes a key and its associated value from the cache.
//
// Parameters:
//   - key: The identifier of the entry to remove
//
// If the key doesn't exist, this operation does nothing. Cache size statistics
// are automatically updated.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	c.updateStatsSize()
}

// Stats returns a snapshot of the current cache statistics.
//
// Returns:
//   - A Stats struct containing hit count, miss count, and current cache size
//
// The returned statistics are consistent at the moment of the call but may
// change immediately after due to concurrent operations.
func (c *Cache) Stats() Stats {
	c.stats.mu.RLock()
	defer c.stats.mu.RUnlock()

	return Stats{
		Hits: c.stats.Hits,
		Miss: c.stats.Miss,
		Size: c.stats.Size,
	}
}

// Len returns the current number of items in the cache.
//
// Returns:
//   - The total count of entries currently stored (including expired ones
//     that haven't been cleaned yet)
func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// Clear removes all entries from the cache, resetting it to an empty state.
// This operation does not affect the TTL or cleanup interval configuration.
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]Item)
	c.updateStatsSize()
}

// GetWithStats retrieves a value and returns it along with current cache
// statistics in a single atomic operation.
//
// Parameters:
//   - key: The identifier of the cached value to retrieve
//
// Returns:
//   - The cached value if found and not expired
//   - A boolean indicating successful retrieval
//   - A snapshot of current cache statistics
//
// This method is useful when both the value and performance metrics are needed.
func (c *Cache) GetWithStats(key string) (interface{}, bool, Stats) {
	value, found := c.Get(key)
	return value, found, c.Stats()
}

// String returns a human-readable representation of the cache state,
// including size and hit rate percentage.
//
// Implements the fmt.Stringer interface.
func (c *Cache) String() string {
	stats := c.Stats()
	hitRate := 0.0
	total := stats.Hits + stats.Miss
	if total > 0 {
		hitRate = float64(stats.Hits) / float64(total) * 100
	}
	return fmt.Sprintf("Cache{size=%d, hits=%d, miss=%d, hitRate=%.1f%%}",
		stats.Size, stats.Hits, stats.Miss, hitRate)
}

// Stop halts the automatic cleanup goroutine. This method is idempotent and
// safe to call multiple times. It should be called when the cache is no longer
// needed to prevent goroutine leaks.
func (c *Cache) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.stopped {
		return
	}

	c.stopped = true
	select {
	case <-c.stopCleanup:
	default:
		close(c.stopCleanup)
	}
}

// startCleanup launches the periodic cleanup goroutine that scans for and
// removes expired entries at configured intervals.
func (c *Cache) startCleanup() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		case <-c.stopCleanup:
			return
		}
	}
}

// cleanup removes all expired entries from the cache. This operation acquires
// a write lock and updates cache size statistics after completion.
func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, item := range c.items {
		if now.After(item.expiration) {
			delete(c.items, key)
		}
	}
	c.updateStatsSize()
}

// incrementHit atomically increments the cache hit counter.
func (c *Cache) incrementHit() {
	c.stats.mu.Lock()
	defer c.stats.mu.Unlock()
	c.stats.Hits++
}

// incrementMiss atomically increments the cache miss counter.
func (c *Cache) incrementMiss() {
	c.stats.mu.Lock()
	defer c.stats.mu.Unlock()
	c.stats.Miss++
}

// updateStatsSize atomically updates the cache size statistic.
func (c *Cache) updateStatsSize() {
	c.stats.mu.Lock()
	defer c.stats.mu.Unlock()
	c.stats.Size = len(c.items)
}
