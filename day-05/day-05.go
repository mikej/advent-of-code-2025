package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("/Users/mike/Downloads/input-day-5.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	rangeStrings, idStrings := SplitInput(string(contents))
	ranges, err := ParseRanges(rangeStrings)
	if err != nil {
		fmt.Println("Error parsing ranges", err)
		return
	}
	ranges = optimizedRanges(ranges)

	ids, err := ParseIds(idStrings)
	if err != nil {
		fmt.Println("Error parsing ids", err)
	}

	count := 0
	for _, id := range ids {
		if isAvailable(id, ranges) {
			count++
		}
	}

	fmt.Printf("Count of items for part 1: %d\n", count)

	availableCount := 0
	for _, r := range ranges {
		availableCount += r.Size()
	}

	fmt.Println("Part 2 total available items:", availableCount)
}

func optimizedRanges(ranges []Range) []Range {
	toDelete := []int{}

outer:
	for i, r := range ranges {
		for j, other := range ranges {
			if i == j || slices.Contains(toDelete, j) {
				continue
			}

			if r.Start >= other.Start && r.End <= other.End {
				toDelete = append(toDelete, i)
				continue outer
			}

			if r.Start <= other.End && r.Start >= other.Start && r.End >= other.End {
				ranges[j].End = r.End
				toDelete = append(toDelete, i)
				continue outer
			}

			if other.End >= r.End && r.End >= other.Start && other.Start >= r.Start {
				ranges[j].Start = r.Start
				toDelete = append(toDelete, i)
				continue outer
			}
		}
	}

	optimised := make([]Range, 0, len(ranges)-len(toDelete))
	for i, r := range ranges {
		if !slices.Contains(toDelete, i) {
			optimised = append(optimised, r)
		}
	}

	return optimised
}

func isAvailable(id int, ranges []Range) bool {
	for _, r := range ranges {
		if r.Contains(id) {
			return true
		}
	}

	return false
}

func SplitInput(input string) ([]string, []string) {
	text := strings.TrimRight(string(input), "\n")

	parts := strings.SplitN(text, "\n\n", 2)
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}

type Range struct {
	Start, End int
}

func (r Range) Contains(n int) bool {
	return n >= r.Start && n <= r.End
}

func (r Range) Size() int {
	return r.End - r.Start + 1
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
