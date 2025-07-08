package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortID(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type URLStore struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewURLStore() *URLStore {
	return &URLStore{
		store: make(map[string]string),
	}
}

func (s *URLStore) Save(shortID, originalURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[shortID] = originalURL
}

func (s *URLStore) Get(shortID string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.store[shortID]
	return url, ok
}

func main() {
	rand.Seed(time.Now().UnixNano())
	store := NewURLStore()

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			URL string `json:"url"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.URL == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		shortID := generateShortID(6)
		store.Save(shortID, req.URL)

		resp := map[string]string{
			"short_url": fmt.Sprintf("http://localhost:8080/%s", shortID),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortID := r.URL.Path[1:]
		if shortID == "" {
			http.Error(w, "No short ID provided", http.StatusBadRequest)
			return
		}
		originalURL, ok := store.Get(shortID)
		if !ok {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusFound)
	})

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
