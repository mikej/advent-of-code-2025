package main

import "testing"

func TestSplitIntoRangesAndAvailableIngredients(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	ranges, availableIds := SplitInput(input)

	if len(ranges) != 4 {
		t.Errorf("Expected 4 ranges, got %d", len(ranges))
	}

	if len(availableIds) != 6 {
		t.Errorf("Expected 6 available ids, got %d", len(availableIds))
	}
}
