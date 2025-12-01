package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/mike/Downloads/input-day-1.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	position := 50
	timesPointedAtZero := 0

	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing distance")
			return
		}

		result := NextPosition(direction, distance, position)
		position = result.nextPosition
		timesPointedAtZero += result.timesPointedAtZero
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}

	fmt.Println("Password is", timesPointedAtZero)
}

type RotationResult struct {
	timesPointedAtZero int
	nextPosition       int
}

func NextPosition(direction string, distance int, currentPosition int) RotationResult {
	change := distance % 100
	completeTurns := distance / 100

	timesPointedAtZero := completeTurns
	var nextPosition int

	if direction == "R" {
		nextPosition = currentPosition + change
	} else if direction == "L" {
		nextPosition = currentPosition - change
	} else {
		panic("Unknown direction")
	}

	if nextPosition >= 100 {
		nextPosition -= 100
		timesPointedAtZero++
	} else if nextPosition < 0 {
		nextPosition = 100 + nextPosition
		if currentPosition != 0 {
			timesPointedAtZero++
		}
	} else if nextPosition == 0 && currentPosition != 0 {
		timesPointedAtZero++
	}
	return RotationResult{timesPointedAtZero, nextPosition}
}
