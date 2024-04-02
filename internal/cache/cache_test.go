package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 10 * time.Millisecond

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://uzee.dev",
			val: []byte("Hello, World!"),
		},
		{
			key: "https://uzeeinc.net",
			val: []byte("Konnichiwa, Sekai!"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)

			// Test adding
			cache.Add(c.key, c.val)

			// Test getting
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key %v in cache", c.key)
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("Expected value %v, got %v", string(c.val), string(val))
				return
			}
		})
	}
}

func TestReap(t *testing.T) {
	const interval = 10 * time.Millisecond
	const waitTime = 2 * interval

	cache := NewCache(interval)

	cache.Add("https://uzee.dev", []byte("Hello, World!"))

	_, ok := cache.Get("https://uzee.dev")
	if !ok {
		t.Errorf("Expected key https://uzee.dev to be in cache")
		return
	}

	time.Sleep(interval)

	cache.Add("https://uzeeinc.net", []byte("Konnichiwa, Sekai!"))

	time.Sleep(interval)

	_, ok = cache.Get("https://uzee.dev")
	if ok {
		t.Errorf("Expected key https://uzee.dev to be reaped")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://uzeeinc.net")
	if ok {
		t.Errorf("Expected key https://uzeeinc.net to be reaped")
		return
	}
}

func TestDelete(t *testing.T) {
	const interval = 10 * time.Millisecond

	cache := NewCache(interval)

	cache.Add("https://uzee.dev", []byte("Hello, World!"))

	_, ok := cache.Get("https://uzee.dev")
	if !ok {
		t.Errorf("Expected key https://uzee.dev to be in cache")
		return
	}

	cache.Delete("https://uzee.dev")

	_, ok = cache.Get("https://uzee.dev")
	if ok {
		t.Errorf("Expected key https://uzee.dev to be deleted")
		return
	}
}

func TestGetNonExistent(t *testing.T) {
	const interval = 10 * time.Millisecond

	cache := NewCache(interval)

	_, ok := cache.Get("https://uzee.dev")
	if ok {
		t.Errorf("Expected key https://uzee.dev to not be in cache")
		return
	}
}

func TestClear(t *testing.T) {
	const interval = 10 * time.Millisecond

	cache := NewCache(interval)

	cache.Add("https://uzee.dev", []byte("Hello, World!"))
	cache.Add("https://uzeeinc.net", []byte("Konnichiwa, Sekai!"))

	cache.Clear()

	_, ok := cache.Get("https://uzee.dev")
	if ok {
		t.Errorf("Expected key https://uzee.dev to be deleted")
		return
	}

	_, ok = cache.Get("https://uzeeinc.net")
	if ok {
		t.Errorf("Expected key https://uzeeinc.net to be deleted")
		return
	}
}

// Test cache concurrency by adding and getting values from multiple goroutines
func TestConcurrentAccess(t *testing.T) {
	const interval = 10 * time.Millisecond
	const numGoroutines = 100

	cache := NewCache(interval)

	done := make(chan struct{})

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			key := fmt.Sprintf("https://uzee.dev/%v", i)
			val := []byte(fmt.Sprintf("Hello, World! %v", i))

			cache.Add(key, val)

			val2, ok := cache.Get(key)
			if !ok {
				t.Errorf("Expected to find key %v in cache", key)
				return
			}

			if string(val) != string(val2) {
				t.Errorf("Expected value %v, got %v", string(val), string(val2))
				return
			}

			done <- struct{}{}
		}(i)
	}

	for i := 0; i < numGoroutines; i++ {
		<-done
	}
}
