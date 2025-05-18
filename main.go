package main

import (
	"log"
	"net/http"

	"github.com/PV6142003/url-shortener/internal/store"
	"github.com/gorilla/mux"
)

func main() {
	urlStore := store.NewURLStore()
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/shorten", urlStore.HandleShorten).Methods("POST")
	r.HandleFunc("/{shortCode}", urlStore.HandleRedirect)

	log.Println("Server running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
