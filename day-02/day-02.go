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
			invalid, err := isInvalid(i)
			if err != nil {
				fmt.Printf("Error checking number %d: %v\n", i, err)
				return
			}
			if invalid {
				total += i
			}
		}
	}

	fmt.Println("Total is", total)
}

func isInvalid(i int64) (bool, error) {
	str := strconv.FormatInt(i, 10)

	for i := 2; i <= len(str); i++ {
		if len(str)%i != 0 {
			continue
		}

		parts, err := parts(str, i)
		if err != nil {
			return false, err
		}

		same, err := allSame(parts)
		if err != nil {
			return false, err
		}
		
		if same {
			return true, nil
		}
	}

	return false, nil
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

func allSame(s []string) (bool, error) {
	if s == nil {
		return false, fmt.Errorf("nil array")
	}
	if len(s) == 0 {
		return true, nil
	}

	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false, nil
		}
	}

	return true, nil
}
