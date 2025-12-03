package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		joltageForLine, err := maxJoltage(line)
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
