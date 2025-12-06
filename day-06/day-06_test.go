package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"123 328 51  64",
		"45  64  387 23",
		"6   98  215 314",
		"*   +   *   +  ",
	}

	result := SolvePart1(input)
	if result != 4277556 {
		t.Errorf("Expected 4277556, got %d", result)
	}
}
