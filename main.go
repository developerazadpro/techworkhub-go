package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Go Matching Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
