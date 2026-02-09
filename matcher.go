package main

import "sort"

func MatchTechnicians(required []string, technicians []Technician) []int {
	var scored []ScoredTech

	for _, tech := range technicians {
		score := 0

		for _, req := range required {
			for _, skill := range tech.Skills {
				if skill == req {
					score++
				}
			}
		}

		if score > 0 {
			scored = append(scored, ScoredTech{
				ID: tech.ID,
				Score: score,
			})
		}
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	var result []int
	for _, tech := range scored {
		result = append(result, tech.ID)
	}

	return result
}