package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)
//
// ------------------DATA STRUCT--------------------------------
//

// Technician represents a technician with skills
type Technician struct {
	ID int `json:"id"`
	Skills []string `json:"skills"`
}

// MatchRequest is the payload received from PHP
type MatchRequest struct {
	JobID int `json:"job_id"`
	RequiredSkills []string `json:"required_skills"`
	Technicians []Technician `json:"technicians"`
}

// MatchResponse is the response sent back to PHP
type MatchResponse struct {
	RecommendedTechnicians []int `json:"recommended_technicians"`
}

// Internal struct for scoring
type ScoredTech struct {
	ID int
	Score int
}

//
// -----------------HANDLERS----------------
//

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// Match technicians based on required skills
func matchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode incoming JSON
	var req MatchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Score technicians
	var scored []ScoredTech 

	for _, tech := range req.Technicians {
		score := 0

		for _, required := range req.RequiredSkills {
			for _, skill := range tech.Skills {
				if skill == required {
					score++
				}
			}
		}

		// Include only technicians with at least one matching skill
		if score > 0 {
			scored = append(scored, ScoredTech{
				ID: tech.ID,
				Score: score,
			})
		}
	}

	// Sort by score (descending)
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	// Prepare response
	var recommended []int 
	for _, tech := range scored {
		recommended = append(recommended, tech.ID)
	}

	resp := MatchResponse{
		RecommendedTechnicians: recommended,
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

//
// ---------- MAIN ----------
//

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/match", matchHandler)

	log.Println("Go Matching Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
