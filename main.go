package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Technician struct {
	ID int `json:"id"`
	Skills []string `json:"skills"`
}

type MatchRequest struct {
	JobID int `json:"job_id"`
	RequiredSkills []string `json:"required_skills"`
	Technicians []Technician `json:"technicians"`
}

type MatchResponse struct {
	RecommendedTechnicians []int `json:"reommended_technicians"`
}

func matchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req MatchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// TEMPORARY: fake recommendation
	var technicianIDs []int 
	for _, tech := range req.Technicians {
		technicianIDs = append(technicianIDs, tech.ID)
	}

	resp := MatchResponse{
		RecommendedTechnicians: technicianIDs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/match", matchHandler)

	log.Println("Go Matching Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
