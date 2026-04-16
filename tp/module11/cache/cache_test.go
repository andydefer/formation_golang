package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Example_basicUsage demonstrates basic cache operations.
func Example_basicUsage() {
	// Create a cache with 5 second TTL and cleanup every 1 second
	c := NewCache(5*time.Second, 1*time.Second)
	defer c.Stop()

	// Store values
	c.Set("username", "alice")
	c.Set("score", 42)

	// Retrieve values
	if value, found := c.Get("username"); found {
		fmt.Printf("Username: %v\n", value)
	}

	// Output: Username: alice
}

// ExampleCache_GetWithStats demonstrates retrieving values with statistics.
func ExampleCache_GetWithStats() {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	c.Set("key", "value")

	value, found, stats := c.GetWithStats("key")
	if found {
		fmt.Printf("Value: %v, Hits: %d\n", value, stats.Hits)
	}

	// Output: Value: value, Hits: 1
}

// ExampleCache_String demonstrates the string representation.
func ExampleCache_String() {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	c.Set("key1", "value1")
	c.Set("key2", "value2")
	c.Get("key1") // Hit

	fmt.Println(c.String())

	// Output: Cache{size=2, hits=1, miss=0, hitRate=100.0%}
}

// TestNewCache verifies cache creation and basic functionality.
func TestNewCache(t *testing.T) {
	c := NewCache(1*time.Second, 100*time.Millisecond)
	defer c.Stop()

	if c == nil {
		t.Fatal("NewCache returned nil")
	}

	if c.Len() != 0 {
		t.Errorf("Expected empty cache, got size %d", c.Len())
	}
}

// TestSetAndGet verifies basic set and get operations.
func TestSetAndGet(t *testing.T) {
	c := NewCache(1*time.Second, 100*time.Millisecond)
	defer c.Stop()

	key := "test-key"
	expected := "test-value"

	c.Set(key, expected)
	actual, found := c.Get(key)

	if !found {
		t.Errorf("Get(%q) returned not found, expected found", key)
	}
	if actual != expected {
		t.Errorf("Get(%q) = %v, expected %v", key, actual, expected)
	}
}

// TestGetNonExistentKey verifies that getting a non-existent key returns not found.
func TestGetNonExistentKey(t *testing.T) {
	c := NewCache(1*time.Second, 100*time.Millisecond)
	defer c.Stop()

	_, found := c.Get("non-existent")

	if found {
		t.Error("Get on non-existent key returned found, expected not found")
	}
}

// TestExpiration verifies that entries expire after TTL.
func TestExpiration(t *testing.T) {
	ttl := 100 * time.Millisecond
	c := NewCache(ttl, 50*time.Millisecond)
	defer c.Stop()

	key := "expiring-key"
	value := "will-expire"

	c.Set(key, value)

	_, found1 := c.Get(key)
	if !found1 {
		t.Error("Entry should exist immediately after Set")
	}

	time.Sleep(ttl + 50*time.Millisecond)

	_, found2 := c.Get(key)
	if found2 {
		t.Error("Entry should have expired and not be found")
	}
}

// TestTTLRefresh verifies that Get refreshes the TTL.
func TestTTLRefresh(t *testing.T) {
	ttl := 200 * time.Millisecond
	c := NewCache(ttl, 50*time.Millisecond)
	defer c.Stop()

	key := "refresh-key"
	value := "will-be-refreshed"

	c.Set(key, value)

	time.Sleep(ttl / 2)

	_, found1 := c.Get(key)
	if !found1 {
		t.Error("Entry should still exist at half TTL")
	}

	time.Sleep(ttl / 2)

	_, found2 := c.Get(key)
	if !found2 {
		t.Error("Entry should exist after TTL refresh")
	}
}

// TestDelete verifies manual deletion of entries.
func TestDelete(t *testing.T) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	key := "delete-key"
	c.Set(key, "value")

	if c.Len() != 1 {
		t.Errorf("Expected size 1, got %d", c.Len())
	}

	c.Delete(key)

	if c.Len() != 0 {
		t.Errorf("Expected size 0 after deletion, got %d", c.Len())
	}

	_, found := c.Get(key)
	if found {
		t.Error("Deleted key should not be found")
	}
}

// TestClear verifies the Clear method removes all entries.
func TestClear(t *testing.T) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	c.Set("key1", "value1")
	c.Set("key2", "value2")
	c.Set("key3", "value3")

	if c.Len() != 3 {
		t.Errorf("Expected size 3, got %d", c.Len())
	}

	c.Clear()

	if c.Len() != 0 {
		t.Errorf("Expected size 0 after Clear, got %d", c.Len())
	}
}

// TestStatistics verifies that hit/miss statistics are correctly tracked.
func TestStatistics(t *testing.T) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	c.Set("existing", "value")

	c.Get("existing")
	c.Get("missing1")
	c.Get("missing2")

	stats := c.Stats()
	if stats.Hits != 1 {
		t.Errorf("Expected 1 hit, got %d", stats.Hits)
	}
	if stats.Miss != 2 {
		t.Errorf("Expected 2 misses, got %d", stats.Miss)
	}
	if stats.Size != 1 {
		t.Errorf("Expected size 1, got %d", stats.Size)
	}
}

// TestConcurrentAccess verifies thread-safety under concurrent read/write operations.
func TestConcurrentAccess(t *testing.T) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	var wg sync.WaitGroup
	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			c.Set(key, id)
		}(i)
	}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			c.Get(key)
		}(i)
	}

	wg.Wait()

	stats := c.Stats()
	if stats.Size == 0 && numGoroutines > 0 {
		t.Logf("Cache size after concurrent access: %d", stats.Size)
	}
}

// TestStopIdempotent verifies that Stop can be called multiple times safely.
func TestStopIdempotent(t *testing.T) {
	c := NewCache(1*time.Second, 100*time.Millisecond)

	c.Stop()
	c.Stop()
	c.Stop()

	time.Sleep(200 * time.Millisecond)
}

// TestLen verifies the Len method returns correct counts.
func TestLen(t *testing.T) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	if c.Len() != 0 {
		t.Errorf("Initial Len() = %d, expected 0", c.Len())
	}

	c.Set("key1", "value1")
	c.Set("key2", "value2")

	if c.Len() != 2 {
		t.Errorf("Len() after adds = %d, expected 2", c.Len())
	}

	c.Delete("key1")
	if c.Len() != 1 {
		t.Errorf("Len() after delete = %d, expected 1", c.Len())
	}
}

// BenchmarkSet measures Set performance.
func BenchmarkSet(b *testing.B) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set("key", i)
	}
}

// BenchmarkGet measures Get performance for existing keys.
func BenchmarkGet(b *testing.B) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	c.Set("key", "value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get("key")
	}
}

// BenchmarkGetMiss measures Get performance for missing keys.
func BenchmarkGetMiss(b *testing.B) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get("missing-key")
	}
}

// BenchmarkConcurrentReadWrite measures performance under concurrent access.
func BenchmarkConcurrentReadWrite(b *testing.B) {
	c := NewCache(1*time.Minute, 30*time.Second)
	defer c.Stop()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := fmt.Sprintf("key-%d", i%100)
			if i%2 == 0 {
				c.Set(key, i)
			} else {
				c.Get(key)
			}
			i++
		}
	})
}
