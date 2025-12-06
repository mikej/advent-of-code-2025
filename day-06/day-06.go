package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/mikej/advent-of-code-2025/shared/input"
)

func main() {

	separator := regexp.MustCompile(`\s+`)

	lines, err := input.ReadFromFile("/Users/mike/Downloads/input-day-6.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	numberLines, operatorLine := lines[:len(lines)-1], lines[len(lines)-1]

	operators := separator.Split(operatorLine, -1)
	calculations := len(operators)

	numbers := make([][]int, len(numberLines))
	for n, line := range numberLines {
		strings := separator.Split(line, -1)
		ints := make([]int, len(strings))
		for i, s := range strings {
			num, _ := strconv.Atoi(s)
			ints[i] = num
		}
		numbers[n] = ints
	}

	grandTotal := 0
	for i := 0; i < calculations; i++ {
		totalSoFar := numbers[0][i]
		fmt.Print(totalSoFar)
		for j := 1; j < len(numbers); j++ {
			if operators[i] == "+" {
				fmt.Printf(" + %d", numbers[j][i])
				totalSoFar += numbers[j][i]
			} else if operators[i] == "*" {
				fmt.Printf(" * %d", numbers[j][i])
				totalSoFar *= numbers[j][i]
			}
		}
		fmt.Printf(" = %d\n", totalSoFar)
		grandTotal += totalSoFar
	}

	fmt.Println("Grand total is", grandTotal)
}
