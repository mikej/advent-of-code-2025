package main

import (
	"fmt"

	"github.com/mikej/advent-of-code-2025/shared/input"
)

func main() {
	lines, err := input.ReadFromFile("/Users/mike/Downloads/input-day-4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	grid := make([][]bool, len(lines))
	for i := range lines {
		grid[i] = make([]bool, len(lines[i]))
	}

	for i := range lines {
		for j := range lines[i] {
			grid[i][j] = string(lines[i][j]) == "@"
		}
	}

	totalRollsRemoved := 0

	for {
		accessibleRolls := 0
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] && countNeighboours(grid, i, j) < 4 {
					grid[i][j] = false
					accessibleRolls++
				}
			}
		}

		if accessibleRolls == 0 {
			break
		}
		totalRollsRemoved += accessibleRolls
	}

	fmt.Println(totalRollsRemoved)
}

func countNeighboours(grid [][]bool, x, y int) int {
	count := 0

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if (i == x && j == y) || i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
				continue
			}
			if grid[i][j] {
				count++
			}
		}
	}

	return count
}
