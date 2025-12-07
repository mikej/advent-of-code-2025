package main

import "testing"

func TestWithExampleData(t *testing.T) {
	input := []string{
		".......S.......",
		"...............",
		".......^.......",
		"...............",
		"......^.^......",
		"...............",
		".....^.^.^.....",
		"...............",
		"....^.^...^....",
		"...............",
		"...^.^...^.^...",
		"...............",
		"..^...^.....^..",
		"...............",
		".^.^.^.^.^...^.",
		"...............",
	}

	manifold, error := NewManifold(input)
	if error != nil {
		t.Errorf("Unexpected error parsing input: %v", error)
		return
	}
	manifold.Run()
	if manifold.SplitCount() != 21 {
		t.Errorf("Expected 21, got %d", manifold.SplitCount())
	}
}
