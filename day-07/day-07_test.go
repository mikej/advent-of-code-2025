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

	manifold := NewManifold(input)
	manifold.Run()
	if manifold.SplitCount() != 21 {
		t.Errorf("Expected 21, got %d", manifold.SplitCount())
	}
}
