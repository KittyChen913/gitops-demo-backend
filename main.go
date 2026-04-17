package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	AppVersion = "1.0.0"
	AppMessage = "hello from backend"
)

var commit = "dev"

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{
		"service": "backend",
		"version": AppVersion,
		"message": AppMessage,
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{
		"status": "ok",
	})
}

func metaHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{
		"service": "backend",
		"version": AppVersion,
		"commit":  commit,
	})
}

func main() {
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/meta", metaHandler)

	log.Printf("GitOps demo backend starting — version %s (commit: %s)", AppVersion, commit)
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
