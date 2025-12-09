package main

import "testing"

var example = []string{
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
}

func TestParseTiles(t *testing.T) {
	tiles, err := parseTiles(example)

	if err != nil {
		t.Errorf("Unexpected error parsing input: %v", err)
	}

	if len(tiles) != 8 {
		t.Errorf("Expected 8 tiles, got %d", len(tiles))
	}
}

func TestLargestArea(t *testing.T) {
	tiles, err := parseTiles(example)

	if err != nil {
		t.Errorf("Unexpected error parsing input: %v", err)
	}

	got := largestArea(tiles)

	if got != 50 {
		t.Errorf("Expected largest area to be 50, got %d", got)
	}
}
