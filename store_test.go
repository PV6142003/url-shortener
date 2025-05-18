package store

import (
	"testing"
	"time"
)

func TestURLStore_SaveAndGet(t *testing.T) {
	store := NewURLStore()
	shortCode := "abc123"
	url := "https://go.dev"

	// Test Save
	store.Save(shortCode, url)

	// Test Get
	if storedURL, exists := store.Get(shortCode); !exists || storedURL != url {
		t.Errorf("Get() = %v, %v; want %v, true", storedURL, exists, url)
	}
}

func TestTrackVisit_Concurrency(t *testing.T) {
	store := NewURLStore()
	shortCode := "xyz789"
	iterations := 1000

	// Concurrent metric tracking
	for i := 0; i < iterations; i++ {
		store.TrackVisit(shortCode)
	}

	// Allow time for goroutines to complete
	time.Sleep(100 * time.Millisecond)

	if count := store.metrics[shortCode]; count != iterations {
		t.Errorf("TrackVisit() count = %v; want %v", count, iterations)
	}
}
