package store

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func (s *URLStore) HandleShorten(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode()
	s.Save(shortCode, request.URL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortCode,
	})
}

func (s *URLStore) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	if url, exists := s.Get(shortCode); exists {
		s.TrackVisit(shortCode) // Async metric tracking
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
	http.NotFound(w, r)
}
