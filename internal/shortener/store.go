package store

import (
	"sync"
)

type URLStore struct {
	urls    map[string]string
	metrics map[string]int
	mutex   sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls:    make(map[string]string),
		metrics: make(map[string]int),
	}
}

func (s *URLStore) Save(shortCode, url string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.urls[shortCode] = url
}

func (s *URLStore) Get(shortCode string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	url, exists := s.urls[shortCode]
	return url, exists
}

// Track visits concurrently
func (s *URLStore) TrackVisit(shortCode string) {
	go func() {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.metrics[shortCode]++
	}()
}
