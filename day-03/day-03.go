package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/mike/Downloads/input-day-3.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Total := 0
	part2Total := 0

	for scanner.Scan() {
		line := scanner.Text()

		joltageForLine, err := maxJoltage2(line, 2)
		if err != nil {
			fmt.Println("Error getting joltage for line", err)
		}
		part1Total += joltageForLine

		joltageForLine, err = maxJoltage2(line, 12)
		if err != nil {
			fmt.Println("Error getting joltage for line", err)
		}
		part2Total += joltageForLine
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	fmt.Println("Part 1 total joltage is", part1Total)
	fmt.Println("Part 2 total joltage is", part2Total)
}

func maxJoltage2(line string, digitsToUse int) (int, error) {
	if len(line) < digitsToUse {
		return 0, fmt.Errorf("need at least %d digits", digitsToUse)
	}

	var startSearchAt = 0
	var digits = make([]string, digitsToUse)

	for digit := 0; digit < digitsToUse; digit++ {
		digits[digit] = string(line[startSearchAt])
		var endSearchAt = len(line) - (digitsToUse - digit)
		var maxIndex = startSearchAt

		for i := startSearchAt + 1; i <= endSearchAt; i++ {
			if string(line[i]) > digits[digit] {
				digits[digit] = string(line[i])
				maxIndex = i
			}
		}

		startSearchAt = maxIndex + 1
	}

	number, _ := strconv.Atoi(strings.Join(digits, ""))
	return number, nil
}
