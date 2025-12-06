package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mikej/advent-of-code-2025/shared/input"
)

func main() {
	lines, err := input.ReadFromFile("/Users/mike/Downloads/input-day-6.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	grandTotal := SolvePart1(lines)

	fmt.Println("Grand total is", grandTotal)
}

func SolvePart1(lines []string) int {
	numberLines, operatorLine := lines[:len(lines)-1], lines[len(lines)-1]

	separator := regexp.MustCompile(`\s+`)

	operators := separator.Split(strings.TrimSpace(operatorLine), -1)
	calculations := len(operators)

	numbers := make([][]int, len(numberLines))
	for n, line := range numberLines {
		numberStrings := separator.Split(line, -1)
		ints := make([]int, len(numberStrings))
		for i, s := range numberStrings {
			num, _ := strconv.Atoi(s)
			ints[i] = num
		}
		numbers[n] = ints
	}

	grandTotal := 0
	for i := 0; i < calculations; i++ {
		totalSoFar := numbers[0][i]
		for j := 1; j < len(numbers); j++ {
			if operators[i] == "+" {
				totalSoFar += numbers[j][i]
			} else if operators[i] == "*" {
				totalSoFar *= numbers[j][i]
			}
		}
		grandTotal += totalSoFar
	}
	return grandTotal
}
