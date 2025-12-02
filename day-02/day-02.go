package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("/Users/mike/Downloads/input-day-2.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	text := strings.TrimRight(string(data), "\n")
	ranges := strings.Split(text, ",")

	var total int64 = 0

	for _, rangeString := range ranges {
		parts := strings.Split(rangeString, "-")

		startRange, _ := strconv.ParseInt(parts[0], 10, 64)
		endRange, _ := strconv.ParseInt(parts[1], 10, 64)

		for i := startRange; i <= endRange; i++ {
			if isInvalid(i) {
				total += i
			}
		}
	}

	fmt.Println("Total is", total)
}

func isInvalid(i int64) bool {
	str := strconv.FormatInt(i, 10)

	if len(str)%2 == 1 {
		return false
	}

	parts, _ := parts(str, 2)
	return allSame(parts)
}

func parts(str string, n int) ([]string, error) {
	totalLength := len(str)

	if totalLength%n != 0 {
		return nil, fmt.Errorf("string length %d is not divisible by %d", totalLength, n)
	}

	partLength := totalLength / n

	var parts = make([]string, n)
	for i := 0; i < n; i++ {
		parts[i] = str[(i * partLength) : (i*partLength)+partLength]
	}

	return parts, nil
}

func allSame(s []string) bool {
	if len(s) == 0 {
		return true
	}

	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}

	return true
}
