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

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		joltageForLine, err := maxJoltage2(line, 12)
		if err != nil {
			fmt.Println("Error getting joltage for line", err)
		}
		total += joltageForLine
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	fmt.Println("Total joltage is", total)
}

func maxJoltage(line string) (int, error) {
	if len(line) < 2 {
		return 0, fmt.Errorf("need at least 2 digits")
	}

	var maxFirstDigit = string(line[0])
	var maxIndex = 0

	for i := 1; i < len(line)-1; i++ {
		if string(line[i]) > maxFirstDigit {
			maxFirstDigit = string(line[i])
			maxIndex = i
		}
	}

	var maxSecondDigit = string(line[maxIndex+1])
	for i := maxIndex + 1; i < len(line); i++ {
		if string(line[i]) > maxSecondDigit {
			maxSecondDigit = string(line[i])
		}
	}

	number, _ := strconv.Atoi(maxFirstDigit + maxSecondDigit)
	return number, nil
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
