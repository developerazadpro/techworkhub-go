package main

import "testing"

func TestMatchTechnicians_Basic(t *testing.T) {
	required := []string{"plumber", "carpenter"}

	technicians := []Technician {
		{ID: 1, Skills: []string{"plumber"}},
		{ID: 2, Skills: []string{"plumber", "carpenter"}},
		{ID: 3, Skills: []string{"plumber", "carpenter", "electrician"}},
	}

	result := MatchTechnicians(required, technicians)

	expected := []int{2, 3, 1}

	if len(result) != len(expected) {
		t.Fatalf("expected %d results, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestMatchTechnicians_SortedByScore(t *testing.T) {
	required := []string{"plumber", "carpenter", "electrician"}

	technicians := []Technician {
		{ID: 1, Skills: []string{"plumber"}},
		{ID: 2, Skills: []string{"plumber", "carpenter"}},
		{ID: 3, Skills: []string{"plumber", "carpenter", "electrician"}},
	}

	result := MatchTechnicians(required, technicians)

	expected := []int{3, 2, 1}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

func TestMatchTechnicians_NoMatches(t *testing.T) {
	required := []string{"electrician"}

	technicians := []Technician {
		{ID: 1, Skills: []string{"plumber"}},
		{ID: 2, Skills: []string{"plumber", "carpenter"}},
	}

	result := MatchTechnicians(required, technicians)

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}