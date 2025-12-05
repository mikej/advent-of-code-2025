package main

import (
	"reflect"
	"testing"
)

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

func TestParseRanges(t *testing.T) {
	result, err := ParseRanges([]string{"3-5", "10-14"})
	if err != nil {
		t.Errorf("Unexpected error parsing ranges: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 ranges, got %d", len(result))
	}
	if result[0].Start != 3 || result[0].End != 5 {
		t.Errorf("Expected range 3-5, got %d-%d", result[0].Start, result[0].End)
	}
	if result[1].Start != 10 || result[1].End != 14 {
		t.Errorf("Expected range 10-14, got %d-%d", result[0].Start, result[0].End)
	}
}

func TestParseAvailableIds(t *testing.T) {
	result, err := ParseIds([]string{"1", "5"})
	if err != nil {
		t.Errorf("Unexpected error parsing ids: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 ids, got %d", len(result))
	}
	if result[0] != 1 {
		t.Errorf("Expected first id to be 1, got %d", result[0])
	}
	if result[1] != 5 {
		t.Errorf("Expected second id to be 5, got %d", result[1])
	}
}

func TestRange_Contains(t *testing.T) {
	r := Range{Start: 3, End: 5}

	if !r.Contains(3) {
		t.Errorf("Expected range 3-5 to contain 3")
	}

	if !r.Contains(4) {
		t.Errorf("Expected range 3-5 to contain 4")
	}

	if !r.Contains(5) {
		t.Errorf("Expected range 3-5 to contain 5")
	}

	if r.Contains(7) {
		t.Errorf("Expected range 3-5 not to contain 7")
	}
}

func TestOptimisedRanges(t *testing.T) {
	tests := []struct {
		name   string
		ranges []Range
		want   []Range
	}{
		{name: "Should be unmodified if ranges don't overlap with each other",
			ranges: []Range{{1, 5}, {8, 12}}, want: []Range{{1, 5}, {8, 12}}},
		{name: "Ranges that are completely covered by another range should be deleted",
			ranges: []Range{{1, 20}, {7, 11}}, want: []Range{{1, 20}}},
		{name: "Contiguous ranges should be combined",
			ranges: []Range{{1, 5}, {5, 10}, Range{10, 15}}, want: []Range{{1, 15}}},
		{name: "Overlapping ranges should be combined", ranges: []Range{{1, 7}, {5, 10}}, want: []Range{{1, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optimizedRanges(tt.ranges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
