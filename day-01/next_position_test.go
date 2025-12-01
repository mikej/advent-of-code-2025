package main

import "testing"

func TestNextPosition(t *testing.T) {
	tests := []struct {
		direction                 string
		distance, currentPosition int
		expected                  int
		pointedAtZero             int
	}{
		{direction: "L", distance: 68, currentPosition: 50, expected: 82, pointedAtZero: 1},
		{direction: "L", distance: 30, currentPosition: 82, expected: 52, pointedAtZero: 0},
		{direction: "R", distance: 48, currentPosition: 52, expected: 0, pointedAtZero: 1},
		{direction: "L", distance: 5, currentPosition: 0, expected: 95, pointedAtZero: 0},
		{direction: "R", distance: 60, currentPosition: 95, expected: 55, pointedAtZero: 1},
		{direction: "L", distance: 55, currentPosition: 55, expected: 0, pointedAtZero: 1},
		{direction: "L", distance: 1, currentPosition: 0, expected: 99, pointedAtZero: 0},
		{direction: "L", distance: 99, currentPosition: 99, expected: 0, pointedAtZero: 1},
		{direction: "R", distance: 14, currentPosition: 0, expected: 14, pointedAtZero: 0},
		{direction: "L", distance: 82, currentPosition: 14, expected: 32, pointedAtZero: 1},
	}

	for _, tt := range tests {
		t.Run("Calculating next position", func(t *testing.T) {
			if got := NextPosition(tt.direction, tt.distance, tt.currentPosition); got.nextPosition != tt.expected || got.timesPointedAtZero != tt.pointedAtZero {
				t.Errorf("NextPosition(%s, %d, %d) = {nextPosition: %d, timesPointedAtZero: %d}; expected {nextPosition: %d, timesPointedAtZero: %d}", tt.direction, tt.distance, tt.currentPosition, got.nextPosition, got.timesPointedAtZero, tt.expected, tt.pointedAtZero)
			}
		})
	}
}
