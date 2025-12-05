package main

import (
	"strconv"
	"strings"
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

type Range struct {
	Start int
	End   int
}

func ParseRanges(ranges []string) ([]Range, error) {
	parsedRanges := make([]Range, len(ranges))
	for i, rangeStr := range ranges {
		parts := strings.Split(rangeStr, "-")
		rangeStart, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		rangeEnd, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		parsedRanges[i] = Range{Start: rangeStart, End: rangeEnd}
	}

	return parsedRanges, nil
}
