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

func TestMaxJoltage2LeavingSome1sAtTheEnd(t *testing.T) {
	checkMaxJoltage2(t, "987654321111111", 987654321111)
}

func TestMaxJoltage2LeavingSome1sInTheMiddle(t *testing.T) {
	checkMaxJoltage2(t, "234234234234278", 434234234278)
}

func TestMaxJoltage2LeavingSomeLowDigits(t *testing.T) {
	checkMaxJoltage2(t, "234234234234278", 434234234278)
}

func TestMaxJoltage2LeavingSome1sNearTheFront(t *testing.T) {
	checkMaxJoltage2(t, "818181911112111", 888911112111)
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

func checkMaxJoltage2(t *testing.T, str string, expected int) {
	result, err := maxJoltage2(str, 12)
	if err != nil {
		t.Errorf("unexpected error getting max joltage: %v", err)
	}
	if result != expected {
		t.Errorf("max joltage should be %d, got %d", expected, result)
	}
}
