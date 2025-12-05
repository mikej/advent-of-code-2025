package main

import (
	"strconv"
	"strings"
)

func SplitInput(input string) ([]string, []string) {
	text := strings.TrimRight(string(input), "\n")

	parts := strings.SplitN(text, "\n\n", 2)
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(n int) bool {
	return n >= r.Start && n <= r.End
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

func ParseIds(strings []string) ([]int, error) {
	ids := make([]int, len(strings))
	for i, str := range strings {
		id, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}
	return ids, nil
}
