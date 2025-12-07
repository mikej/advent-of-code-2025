package main

import "testing"

var example = []string{
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

func TestPart1(t *testing.T) {
	manifold, error := NewManifold(example)
	if error != nil {
		t.Errorf("Unexpected error parsing input: %v", error)
		return
	}
	manifold.Run()
	if manifold.SplitCount() != 21 {
		t.Errorf("Expected 21, got %d", manifold.SplitCount())
	}
}

func TestPart2(t *testing.T) {
	manifold, error := NewManifold(example)
	if error != nil {
		t.Errorf("Unexpected error parsing input: %v", error)
		return
	}
	manifold.Run()

	got := manifold.WorldCount()
	if got != 40 {
		t.Errorf("Expected 40, got %d", got)
	}
}
