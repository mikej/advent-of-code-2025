package main

import (
	"strconv"
	"strings"
)

func SplitInput(input string) ([]string, []string) {
	parts := strings.SplitN(input, "\n\n", 2)
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
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
