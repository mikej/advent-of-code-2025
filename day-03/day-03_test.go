package main

import "testing"

func TestMaxJoltageWhereDigitsAreTogether(t *testing.T) {
	checkMaxJoltage(t, "987654321111111", 98)
}

func TestMaxJoltageWhereDigitsAreAtEachEnd(t *testing.T) {
	checkMaxJoltage(t, "811111111111119", 89)
}

func TestMaxJoltageWhereDigitsAreSeparatedPartWayThrough(t *testing.T) {
	checkMaxJoltage(t, "818181911112111", 92)
}

func TestMaxJoltageReturnsAnErrorWhenStringIsTooShort(t *testing.T) {
	result, err := maxJoltage("1")
	if err == nil {
		t.Errorf("expected to get an error, got no error and a result of %d", result)
	}
}

func checkMaxJoltage(t *testing.T, str string, expected int) {
	result, err := maxJoltage(str)
	if err != nil {
		t.Errorf("unexpected error getting max joltage: %v", err)
	}
	if result != expected {
		t.Errorf("max joltage should be %d, got %d", expected, result)
	}
}
